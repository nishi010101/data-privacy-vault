package encryption

import (
	"crypto/aes"
	"crypto/cipher"
)

func Decrypt(cipherText []byte, nonce []byte, encryptionKey string) (string, error) {
	block, err := aes.NewCipher([]byte(encryptionKey))
	if err != nil {
		panic(err.Error())
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	value, err := aesgcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return "", err
	}
	return string(value), nil
}
