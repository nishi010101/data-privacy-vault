package core

import (
	"context"
	"crypto/aes"
	_ "crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/json"
	"github.com/google/uuid"
	"io"
	"vault/internal/config"
	"vault/internal/model"
)

func Tokenize(context context.Context, values map[string]string, config config.Config) (*map[string]string, error) {
	tokenizedValues := make(map[string]string)
	block, err := aes.NewCipher([]byte(config.EncryptionKey))
	if err != nil {
		panic(err.Error())
	}

	for key, val := range values {
		token := uuid.New().String()
		nonce := make([]byte, 12)
		if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
			panic(err.Error())
		}
		aesgcm, err := cipher.NewGCM(block)
		if err != nil {
			panic(err.Error())
		}
		ciphertext := aesgcm.Seal(nil, nonce, []byte(val), nil)
		marshal, err := json.Marshal(model.TokenizedValue{
			Nonce:           nonce,
			CipherTextValue: ciphertext,
			Token:           token,
		})
		if err != nil {
			return nil, err
		}
		err = config.RedisClient.Set(context, key, marshal, 0).Err()
		if err != nil {
			return nil, err
		}
		tokenizedValues[key] = token
	}
	return &tokenizedValues, nil
}
