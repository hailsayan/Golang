package main

import "fmt"

func main() {
	var str = struct {
		firstName string
		lastName  string
	}{"John", "Doe"}

	fmt.Printf("anonymous struct: %+v\n", str)
}
