package seveneleven

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/lynnau/fuelfor.cheap-api/des"
)

func generateKey() string {
	a := []int{103, 180, 267, 204, 390, 504, 497, 784, 1035, 520, 1155, 648, 988, 1456, 1785}
	b := []int{50, 114, 327, 276, 525, 522, 371, 904, 1017, 810, 858, 852, 1274, 1148, 915}
	c := []int{74, 220, 249, 416, 430, 726, 840, 568, 1017, 700, 1155, 912, 1118, 1372}

	length := len(a) + len(b) + len(c)
	var key string

	for i := 0; i < length; i++ {
		if i%3 == 0 {
			key += fmt.Sprintf("%c", rune(a[i/3]/((i/3)+1)))
		}
		if i%3 == 1 {
			key += fmt.Sprintf("%c", rune(b[(i-1)/3]/(((i-1)/3)+1)))
		}
		if i%3 == 2 {
			key += fmt.Sprintf("%c", rune(c[(i-1)/3]/(((i-2)/3)+1)))
		}
	}

	return key
}

func desDecryptString(deviceID string) string {
	// we only need the first 8 characters of the encryption key
	key := "co.vmob.sdk.android.encrypt.key"[:8]
	prefix := "co.vmob.android.sdk."

	// now to actually encrypt it?
	encrypted, err := des.DesEncrypt([]byte(fmt.Sprintf("%s%s", prefix, deviceID)), []byte(key))
	if err != nil {
		panic(err)
	}

	// return the encrypted message back, base64 encoded
	encoded := base64.StdEncoding.EncodeToString(encrypted)
	encoded = strings.ReplaceAll(encoded, "/", "_")

	// the encoded string must be suffixed with an underscore
	return fmt.Sprintf("%s_", encoded)
}
