package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/hello", handleHelloWorld)
	http.HandleFunc("/getProduct/", getProduct)
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

func getProduct(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	productID := parts[2]
	fmt.Fprintf(w, "Product ID: %s\n", productID)
}
