package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func JSONResponse(w http.ResponseWriter, data interface{}) {
	body, err := json.Marshal(data)

	if err != nil {
		log.Printf("JSON Response fails: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(body)

	if err != nil {
		log.Printf("Fail on response: %v", err)
		return
	}
}
