package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(data)
}

func RespondWithError(w http.ResponseWriter, statusCode int, message string) {
	if statusCode > 499 {
		log.Printf("Responding with 5XX error: %s", message)
	}
	type errorResponse struct {
		Error string `json:"error"`
	}
	RespondWithJSON(w, statusCode, errorResponse{
		Error: message,
	})
}
