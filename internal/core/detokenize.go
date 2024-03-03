package core

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/json"
	"vault/internal/config"
	"vault/internal/model"
)

func Detokenize(context context.Context, values map[string]string, config config.Config) (*map[string]model.DetokenizedValue, error) {
	detokenizedValues := make(map[string]model.DetokenizedValue)
	block, err := aes.NewCipher([]byte(config.EncryptionKey))
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	for key, token := range values {
		value, err := config.RedisClient.Get(context, key).Result()
		if err != nil {
			return nil, err
		}

		var persistedValue model.TokenizedValue
		err = json.Unmarshal([]byte(value), &persistedValue)
		if err != nil {
			return nil, err
		}

		plaintext, err := aesgcm.Open(nil, persistedValue.Nonce, persistedValue.CipherTextValue, nil)
		if err != nil {
			panic(err.Error())
		}

		if persistedValue.Token == token {
			detokenizedValues[key] = model.DetokenizedValue{
				Found: true,
				Value: string(plaintext),
			}
		} else {
			detokenizedValues[key] = model.DetokenizedValue{
				Found: false,
				Value: "",
			}
		}
	}
	return &detokenizedValues, nil
}
