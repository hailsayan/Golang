package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string `json:"Fname"`
	Age   int    `json:"Age"`
	Email string `json:"Mail"`
}

func main() {
	person := Person{
		Name:  "John Doe",
		Age:   21,
		Email: "john.doe@example.com",
	}
	jsonData, err := json.Marshal(person)
	if err != nil {
		panic("failed to marshal")
	}
	jsonString := string(jsonData)
	fmt.Println(jsonString)
}
