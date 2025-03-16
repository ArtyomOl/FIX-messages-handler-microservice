package main

import (
	"FIX-messages-handler-API/server"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "fix api")
	})
	api.HandleFunc("/send", server.GetFixMessage()).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":8080", r))
}
