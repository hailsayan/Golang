package main

import "fmt"

func main() {
	var fruits = map[string]int{
		"apple":  1,
		"banana": 2,
		"orange": 3,
	}

	for i, v := range fruits {
		fmt.Printf("Key : %v, Value : %v\n", i, v)
	}
}
