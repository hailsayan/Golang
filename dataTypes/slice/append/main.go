package main

import "fmt"

func main() {
	letters := []string{"a", "b", "c"}
	letters = append(letters, "e")
	fmt.Println("slice:", letters)

	letters = append(letters, "f", "g", "h")
	fmt.Println("slice:", letters)
}
