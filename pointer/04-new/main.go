package main

import "fmt"

func main() {
	p := new(string)
	*p = "newPointer"
	fmt.Println(*p)
	fmt.Println(p)
}
