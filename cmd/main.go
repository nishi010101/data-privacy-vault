package main

import (
	"encoding/json"
	mux2 "github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
	"net/http"
	"os"
	"vault/internal/config"
	"vault/internal/core"
	"vault/internal/model"
)

func main() {
	redisHost := os.Getenv("REDIS_HOST_URL")
	encryptionKey := os.Getenv("DATA_ENCRYPTION_KEY")
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	appConfig := config.Config{RedisClient: rdb, EncryptionKey: encryptionKey, TokenizeApiKey: "test-tokenize-api-key", DetokenizeApiKey: "test-detokenize-api-key"}
	mux := mux2.NewRouter()
	mux.HandleFunc("/tokenize", tokenizeHandler(appConfig)).Methods(http.MethodPost)
	mux.HandleFunc("/detokenize", detokenizedHandler(appConfig)).Methods(http.MethodPost)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		return
	}
}

func tokenizeHandler(config config.Config) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		headers := request.Header
		apiKey, exists := headers["Vault-Api-Key"]
		if !exists || apiKey[0] != config.TokenizeApiKey {
			http.Error(writer, "Unauthorized request", http.StatusUnauthorized)
			return
		}
		var data model.Request
		err := json.NewDecoder(request.Body).Decode(&data)
		tokenizedData, err := core.Tokenize(request.Context(), data.Data, config)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		response := model.TokenizeResponse{
			Id:   data.Id,
			Data: *tokenizedData,
		}
		err = json.NewEncoder(writer).Encode(response)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
	}
}

func detokenizedHandler(config config.Config) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		headers := request.Header
		apiKey, exists := headers["Vault-Api-Key"]
		if !exists || apiKey[0] != config.DetokenizeApiKey {
			http.Error(writer, "Unauthorized request", http.StatusUnauthorized)
			return
		}

		var data model.Request
		err := json.NewDecoder(request.Body).Decode(&data)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		detokenizedData, err := core.Detokenize(request.Context(), data.Data, config)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		err = json.NewEncoder(writer).Encode(detokenizedData)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
