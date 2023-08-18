package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

var bytesData = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

func Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func Decode(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

var Encrypt = func(text, MySecret string) (string, error) {
	if len(MySecret) < 16 {
		for i := 0; i < 16-len(MySecret); i++ {
			MySecret += "."
		}
	}
	block, err := aes.NewCipher([]byte(MySecret))
	if err != nil {
		return "", err
	}
	plainText := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, bytesData)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	return Encode(cipherText), nil
}

var Decrypt = func(text, MySecret string) (string, error) {
	if len(MySecret) < 16 {
		for i := 0; i < 16-len(MySecret); i++ {
			MySecret += "."
		}
	}
	block, err := aes.NewCipher([]byte(MySecret))
	if err != nil {
		return "", err
	}
	cipherText := Decode(text)
	cfb := cipher.NewCFBDecrypter(block, bytesData)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)
	return string(plainText), nil
}
