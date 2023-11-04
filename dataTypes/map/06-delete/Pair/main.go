package main

import "fmt"

func main() {
	var fruits = map[string]int{
		"apple":  1,
		"banana": 2,
		"orange": 3,
	}
	fmt.Println(len(fruits))

	for key := range fruits {
		delete(fruits, key)
	}
	fmt.Println(len(fruits))
}
