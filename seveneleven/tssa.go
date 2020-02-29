package seveneleven

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

var magic = "yvktroj08t9jltr3ze0isf7r4wygb39s"

func generateTSSA(host, method, payload, accessToken string) (string, error) {
	uri := strings.ToLower(strings.ReplaceAll(host, "https", "http"))
	timestamp := time.Now().Unix()
	id := uuid.New().String()
	mystery := fmt.Sprintf("%s%s%s%d%s", magic, method, uri, timestamp, id)

	if payload != "" {
		// md5 the payload
		checksum := md5.Sum([]byte(payload))
		// and then base64 the md5 checksum/digest
		b64 := base64.StdEncoding.EncodeToString(checksum[:])
		mystery += b64
	}

	mac := hmac.New(sha256.New, encryptionKey)
	mac.Write([]byte(mystery))
	signature := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	tssa := fmt.Sprintf("tssa %s:%s:%s:%d", magic, signature, id, timestamp)

	if accessToken != "" {
		tssa += fmt.Sprintf(":%s", accessToken)
	}

	return tssa, nil
}
