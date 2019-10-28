package encrypt

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
)

func SHA256(source string, secretKey string) []byte {
	key := []byte(secretKey)
	mac := hmac.New(sha256.New, key)
	mac.Write([]byte(source))
	return mac.Sum(nil)
}

func SHA1(source string) string {
	sha1 := sha1.New()
	sha1.Write([]byte(source))
	return hex.EncodeToString(sha1.Sum([]byte(nil)))
}
