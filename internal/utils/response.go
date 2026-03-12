package utils

import (
	"encoding/json"
	"net/http"

	"github.com/hogiabao7725/blog-rest-api-golang/internal/dto/response"
)

func WriteJSON(w http.ResponseWriter, statusCode int, payload response.Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(payload); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}

func WriteError(w http.ResponseWriter, statusCode int, errMsg string) {
	WriteJSON(w, statusCode, response.Response{
		Success: false,
		Error:   errMsg,
	})
}
