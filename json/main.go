package main

import (
	"encoding/json"
	"fmt"
)

// func Marshal(v interface{}) ([]byte, error)

func main() {
	j := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	bytes, err := json.Marshal(j)
	if err != nil {
		panic("failed to marshal")
	}
	fmt.Println(string(bytes))
}
