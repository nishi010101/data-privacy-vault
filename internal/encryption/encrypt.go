package encryption

import (
	"crypto/aes"
	"crypto/cipher"
)

func Encrypt(value string, nonce[]byte, encryptionKey string) []byte  {
	block, err := aes.NewCipher([]byte(encryptionKey))
	if err != nil {
		panic(err.Error())
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	return aesgcm.Seal(nil, nonce, []byte(value), nil)
}
