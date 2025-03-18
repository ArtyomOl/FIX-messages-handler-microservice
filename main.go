package main

import (
	"FIX-messages-handler-API/server"
	"FIX-messages-handler-API/storage"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	redis_cfg := storage.Config{
		Addr:        "localhost:6379",
		DB:          0,
		MaxRetries:  5,
		DialTimeout: 10 * time.Second,
		Timeout:     5 * time.Second,
	}

	redis_client, err := storage.NewClient(context.Background(), redis_cfg)
	if err != nil {
		panic(err)
	}

	go server.StartConsumer(redis_client)

	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "fix api")
	})
	api.HandleFunc("/getorderbook", server.GetFixMessage(redis_client)).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":8080", r))
}

// curl localhost:8080/api/getorderbook --data "{\"symbol\":\"AAPL\",\"depth\":1}"
