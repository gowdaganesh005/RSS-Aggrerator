package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func Err_Response(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Responding with 5XX error:", msg)

	}
	type err_resp struct {
		Error string `json:"error"`
	}

	JSON_Response(w, 400, err_resp{
		Error: msg,
	})
}

func JSON_Response(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal the Json response :%v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
