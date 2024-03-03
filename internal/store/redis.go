package store

import (
	"context"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"vault/internal/model"
)

func PersistData(context context.Context, values map[string]string, redisClient redis.Client) map[string]string {
	tokenizedValues := make(map[string]string)
	for key, val := range values {
		token := uuid.New().String()
		redisClient.Set(context, key, model.TokenizedValue{
			CipherTextValue: val,
			Token:           token,
		}, 0)
		tokenizedValues[key] = token
	}
	return tokenizedValues
}
