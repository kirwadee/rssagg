package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Responding with 5xx error: ", msg)
	}

	type errResponse struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, code, errResponse{Error: msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	//Marshal the payload into JSON string object
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("failed to marshal JSON response:%v", payload)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//add header to the response that we are responding with JSON
	w.Header().Add("Content_Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
