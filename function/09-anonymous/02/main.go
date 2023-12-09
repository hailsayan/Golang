package main

import "fmt"

func main() {
	sum := func(a, b int) {
		fmt.Println(a + b)
	}
	sum(40, 60)
	sum(35, 75)
}
