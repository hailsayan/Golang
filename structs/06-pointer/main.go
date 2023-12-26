package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	p := Person{FirstName: "John", LastName: "Doe", Age: 25}
	ptr := &p
	fmt.Println(ptr.FirstName)
	ptr.FirstName = "lucas"
	fmt.Printf(p.FirstName)
}
