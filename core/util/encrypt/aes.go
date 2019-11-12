package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

func AesDecrypt(key string, encrypt string) (string, error){
	kbs := SHA256(key)
	decode, err := base64.StdEncoding.DecodeString(encrypt)
	if err != nil {
		return "", err
	}
	if len(decode) < aes.BlockSize {
		return "", errors.New("密文太短啦")
	}
	iv := decode[:aes.BlockSize]
	block, err := aes.NewCipher(kbs)
	if err != nil{
		return "", err
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	plantText := make([]byte, len(decode))
	blockMode.CryptBlocks(plantText, decode)
	plantText = PKCS7UnPadding(plantText)
	plantText = plantText[aes.BlockSize:]
	return string(plantText), nil
}

func PKCS7UnPadding(plantText []byte) []byte {
	length := len(plantText)
	unpadding := int(plantText[length-1])
	return plantText[:(length - unpadding)]
}