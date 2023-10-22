package main

import "fmt"

func main() {
	var fruits = map[string]int{
		"apple":  2,
		"banana": 5,
		"orange": 8,
	}
	fmt.Println(fruits["apple"])
	fruit["apple"] = 8
	fmt.Println(fruits["apple"])
}
