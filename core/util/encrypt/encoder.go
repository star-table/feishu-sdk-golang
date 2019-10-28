package encrypt
import (
   "encoding/base64"
	"net/url"
)

func BASE64(input []byte) string{
	return base64.StdEncoding.EncodeToString(input)
}

func URLEncode(input string) string{
	return url.QueryEscape(input)
}