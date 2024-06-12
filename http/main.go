package main

import (
	"fmt"
	"net/http"
)

func welcomePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome"))
}

func main() {
	http.HandleFunc("/welcome", welcomePage)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
