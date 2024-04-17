package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("First Defer statement")
	defer fmt.Println("Second Defer statement")
	panic("An error occurred")
}
