package main

import "fmt"

func main() {
	var fruits = map[string]int{
		"apple":  2,
		"banana": 5,
		"orange": 8,
	}

	mango, ok := fruits["mango"]
	fmt.Printf("Key : %s, Value : %d, Present : %v\n", "mango", mango, ok)

	apple, ok := fruits["apple"]
	fmt.Printf("Key : %s, Value : %d, Present : %v", "apple", apple, ok)
}
