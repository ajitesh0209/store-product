package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Code    int
	Payload interface{}
}

type Fields map[string]interface{}

func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err := w.Write(response)
	if err != nil {
		log.Fatal("Error writing response to JSON", err.Error())
	}
}
