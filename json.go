package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, status int, message string) {
	if status > 499 {
		log.Println("Responding with 5XX error:", message)
	}

	type errorResponse struct {
		Error string `json:"error"`
	}

	respondWithJson(w, status, errorResponse{Error: message})
}

func respondWithJson(w http.ResponseWriter, status int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Println("Failed to marshal JSON response:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(data)
	log.Println("Response sent with status:", status)
}
