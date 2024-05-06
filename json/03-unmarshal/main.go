package main

import (
	"encoding/json"
	"fmt"
)

type Computer struct {
	CPU  string `json:"cpu"`
	RAM  string `json:"ram"`
	Disk string `json:"disk"`
}

func main() {
	jsonData := []byte(`{
		"cpu": "intel",
		"ram": "16GB",
		"disk": "512GB"
	}`)
	var com Computer
	err := json.Unmarshal(jsonData, &com)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("cpu:", com.CPU)
	fmt.Println("ram:", com.RAM)
	fmt.Println("disk:", com.Disk)
}
