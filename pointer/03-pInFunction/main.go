package main

import "fmt"

func exampleFunc(p *int) int {
	return *p
}

func main() {
	p := 10
	fmt.Println(exampleFunc(&p))
}
