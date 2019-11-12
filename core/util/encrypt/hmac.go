package encrypt

import (
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
)

func SHA256(source string) []byte {
	mac := sha256.New()
	mac.Write([]byte(source))
	return mac.Sum(nil)
}

func SHA1(source string) string {
	sha1 := sha1.New()
	sha1.Write([]byte(source))
	return hex.EncodeToString(sha1.Sum([]byte(nil)))
}
