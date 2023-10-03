package main

import "fmt"

func main() {
	letters := []string{"a", "b", "c", "d"}
	// method 1
	for _, letter := range letters {
		fmt.Printf("%s\t", letter)
	}
	fmt.Printf("\n")
	// method2
	for i := 0; i < len(letters); i++ {
		fmt.Printf("%s\t", letters[i])
	}
}
