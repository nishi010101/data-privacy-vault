package core

import (
	"context"
	"encoding/json"
	"vault/internal/config"
	"vault/internal/encryption"
	"vault/internal/model"
)

func Detokenize(context context.Context, values map[string]string, config config.Config) (*map[string]model.DetokenizedValue, error) {
	detokenizedValues := make(map[string]model.DetokenizedValue)

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

		plaintext, err := encryption.Decrypt(persistedValue.CipherTextValue, persistedValue.Nonce, config.EncryptionKey)
		if err != nil {
			return nil, err
		}

		if persistedValue.Token == token {
			detokenizedValues[key] = model.DetokenizedValue{
				Found: true,
				Value: plaintext,
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
