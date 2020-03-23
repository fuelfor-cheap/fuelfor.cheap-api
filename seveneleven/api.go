package seveneleven

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

var (
	// hardcoded constants required for making requests to the 7/11 api
	uri        = "https://711-goodcall.api.tigerspike.com/api/v1"
	deviceName = "SM-G973FZKAXSA"
	// appVersion     = "1.8.0.2027"
	appVersion     = "1.10.2044"
	androidVersion = "Android 9.0.0"

	encryptionKey []byte
)

// Login ...
func Login(email, password, accessToken, deviceID string) (*LoginResponse, string, string) {
	// create the payload
	payload := fmt.Sprintf("{\"Email\":\"%s\",\"Password\":\"%s\",\"DeviceName\":\"%s\",\"DeviceOsNameVersion\":\"%s\"}", email, password, deviceName, androidVersion)

	if deviceID == "" {
		// generate a unique device id
		characters := strings.Split("0123456789abcdef", "")
		for i := 0; i < 16; i++ {
			deviceID += characters[rand.Intn(len(characters)-0)+0]
		}
	}

	url := uri + "/account/login"
	method := "POST"
	tssa, err := generateTSSA(url, method, payload, accessToken)
	if err != nil {
		fmt.Println("err", err)
		panic(err)
	}

	// send a request to the 7/11 api now with our setup headers
	req, err := http.NewRequest(method, url, bytes.NewBuffer([]byte(payload)))
	if err != nil {
		panic(err)
	}
	// add all the headers required for authentication
	req.Header.Add("User-Agent", "Apache-HttpClient/UNAVAILABLE (java 1.4)")
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	req.Header.Add("Connection", "Keep-Alive")
	req.Header.Add("Host", "711-goodcall.api.tigerspike.com")
	req.Header.Add("Authorization", tssa)
	req.Header.Add("X-OsVersion", androidVersion)
	req.Header.Add("X-OsName", "Android")
	req.Header.Add("X-DeviceID", deviceID)
	req.Header.Add("X-AppVersion", appVersion)

	// send the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// read the body and pipe it into a known struct
	bits, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	loginResp := &LoginResponse{}
	if err = json.Unmarshal(bits, loginResp); err != nil {
		panic(err)
	}

	return loginResp, resp.Header.Get("X-Accesstoken"), deviceID
}

// Lock ...
func Lock(latitude, longitude float64, accessToken, deviceSecret, deviceID, fuelType, accountID string) *LockResponse {
	// timestamp, seconds since epoch
	now := time.Now().Unix()

	// create the payload
	payload := fmt.Sprintf("{\"LastStoreUpdateTimestamp\":\"%d\",\"Latitude\":\"%f\",\"Longitude\":\"%f\"}", now, latitude, longitude)

	url := uri + "/FuelLock/StartSession"
	method := "POST"
	tssa, err := generateTSSA(url, method, payload, accessToken)
	if err != nil {
		fmt.Println("err", err)
		panic(err)
	}

	// send a request to the 7/11 api now with our setup headers
	req, err := http.NewRequest(method, url, bytes.NewBuffer([]byte(payload)))
	if err != nil {
		panic(err)
	}

	// decrypt our device id into a "mob id"?
	vMobID := desDecryptString(deviceID)

	req.Header.Add("User-Agent", "Apache-HttpClient/UNAVAILABLE (java 1.4)")
	req.Header.Add("Authorization", tssa)
	req.Header.Add("X-OsVersion", androidVersion)
	req.Header.Add("X-OsName", "Android")
	req.Header.Add("X-DeviceID", deviceID)
	req.Header.Add("X-VmobID", vMobID)
	req.Header.Add("X-AppVersion", appVersion)
	req.Header.Add("X-DeviceSecret", deviceSecret)
	req.Header.Add("Content-Type", "application/json; charset=utf-8")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bits, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	lockSessResp := &LockSessionResponse{}
	if err = json.Unmarshal(bits, lockSessResp); err != nil {
		panic(err)
	}

	matchingEan := fuelTypeToEan[fuelType]
	litresToLock := 150
	payload = fmt.Sprintf("{\"AccountId\":\"%s\",\"FuelType\":\"%d\",\"NumberOfLitres\":\"%d\"}", accountID, matchingEan, litresToLock)

	url = uri + "/FuelLock/Confirm"
	tssa, err = generateTSSA(url, method, payload, accessToken)
	if err != nil {
		panic(err)
	}

	// send a request to the 7/11 api now with our setup headers
	req, err = http.NewRequest(method, url, bytes.NewBuffer([]byte(payload)))
	if err != nil {
		panic(err)
	}

	req.Header.Add("User-Agent", "Apache-HttpClient/UNAVAILABLE (java 1.4)")
	req.Header.Add("Authorization", tssa)
	req.Header.Add("X-OsVersion", androidVersion)
	req.Header.Add("X-OsName", "Android")
	req.Header.Add("X-DeviceID", deviceID)
	req.Header.Add("X-VmobID", vMobID)
	req.Header.Add("X-AppVersion", appVersion)
	req.Header.Add("X-DeviceSecret", deviceSecret)
	req.Header.Add("Content-Type", "application/json; charset=utf-8")

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bits, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	lockResp := &LockResponse{}
	if err = json.Unmarshal(bits, lockResp); err != nil {
		panic(err)
	}
	return lockResp
}

// GetLock ...
func GetLock(accessToken, deviceSecret, deviceID string) []*LockResponse {
	url := uri + "/FuelLock/List"
	method := "GET"
	tssa, err := generateTSSA(url, method, "", accessToken)
	if err != nil {
		fmt.Println("err", err)
		panic(err)
	}

	// send a request to the 7/11 api now with our setup headers
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("User-Agent", "Apache-HttpClient/UNAVAILABLE (java 1.4)")
	req.Header.Add("Authorization", tssa)
	req.Header.Add("X-OsVersion", androidVersion)
	req.Header.Add("X-OsName", "Android")
	req.Header.Add("X-DeviceID", deviceID)
	req.Header.Add("X-AppVersion", appVersion)
	req.Header.Add("X-DeviceSecret", deviceSecret)
	req.Header.Add("Content-Type", "application/json; charset=utf-8")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bits, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	lockResp := []*LockResponse{}
	if err = json.Unmarshal(bits, &lockResp); err != nil {
		panic(err)
	}
	return lockResp
}

func init() {
	// generate our encryption key on startup
	// we'll need this when creating the tssa values
	var err error
	key := generateKey()
	encryptionKey, err = base64.StdEncoding.DecodeString(key)
	if err != nil {
		panic(err)
	}
}
