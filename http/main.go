package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	ID    int     `json:"ID"`
	Name  string  `json:"Name"`
	Price float64 `json:"Price"`
	Count int     `json:"Count"`
}

func welcomePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to Store!"))
}
func allProducts(w http.ResponseWriter, r *http.Request) {
	sampleProducts := []Product{
		{ID: 1, Name: "Paper", Price: 1000, Count: 300},
		{ID: 2, Name: "Pencil", Price: 5000, Count: 1000},
		{ID: 3, Name: "Marker", Price: 20000, Count: 5},
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sampleProducts)

	headers := r.Header
	fmt.Println("Request Headers:")
	for key, values := range headers {
		fmt.Printf("%s: %v\n", key, values)
	}
	method := r.Method
	path := r.URL.Path
	fmt.Printf("HTTP Method: %s\n", method)
	fmt.Printf("Path: %s\n", path)
}
func main() {
	http.HandleFunc("/welcome", welcomePage)
	http.HandleFunc("/products", allProducts)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
