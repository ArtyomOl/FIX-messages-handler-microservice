package server

import (
	"FIX-messages-handler-API/storage"
	"encoding/json"
	"net/http"

	"github.com/redis/go-redis/v9"
)

func GetFixMessage(client *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Symbol string `json:"symbol"`
			Depth  int    `json:"depth"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid request", http.StatusBadRequest)
			return
		}

		orderBook, err := storage.GetOrderBook(client, req.Symbol, req.Depth)
		if err != nil {
			panic(err)
		}
		b, _ := json.Marshal(orderBook)

		w.Write(b)
	}
}
