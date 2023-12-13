package main

import "fmt"

func main() {
	defer fmt.Println("forth")
	defer fmt.Println("third")
	defer fmt.Println("second")

	fmt.Println("first")
}
