package config

import "github.com/redis/go-redis/v9"

type Config struct {
	RedisClient      *redis.Client
	EncryptionKey    string
	TokenizeApiKey   string
	DetokenizeApiKey string
}
