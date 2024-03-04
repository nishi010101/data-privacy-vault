package core

import (
	"context"
	_ "crypto/aes"
	"crypto/rand"
	"encoding/json"
	"github.com/google/uuid"
	"io"
	"vault/internal/config"
	"vault/internal/encryption"
	"vault/internal/model"
)

func Tokenize(context context.Context, values map[string]string, config config.Config) (*map[string]string, error) {
	tokenizedValues := make(map[string]string)
	for key, val := range values {
		//TODO: use something else apart from uuid
		token := uuid.New().String()
		nonce := make([]byte, 12)
		if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
			panic(err.Error())
		}
		ciphertext := encryption.Encrypt(val, nonce, config.EncryptionKey)
		marshal, err := json.Marshal(model.TokenizedValue{
			Nonce:           nonce,
			CipherTextValue: ciphertext,
			Token:           token, //TODO: encrpty this also
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
