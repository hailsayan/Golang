package main

import "fmt"

func main() {
	defer fmt.Println("second")
	defer fmt.Println("forth")
	defer fmt.Println("third")

	fmt.Println("first")
}
