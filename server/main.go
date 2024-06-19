package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", handleHelloWorld)
	http.ListenAndServe(":8080", nil)
}

func handleHelloWorld(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"message": "hello world",
		"status":  "success",
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
