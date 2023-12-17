package main

import "fmt"

func main() {
	var pi float64 = 3.14
	var P *float64 = &pi

	fmt.Println(P)
	fmt.Println(*P)
}
