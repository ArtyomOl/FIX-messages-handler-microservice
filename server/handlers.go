package server

import (
	"FIX-messages-handler-API/fix"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetFixMessage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct{ Message string }
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid request", http.StatusBadRequest)
			return
		}
		encode_message, _ := fix.ParseFixMessages(req.Message)
		fmt.Println(encode_message)
	}
}
