package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func increamentAge(p *Person) {
	p.Age++
}

func main() {
	p := Person{FirstName: "John", LastName: "Doe", Age: 25}
	ptr := &p
	fmt.Println(ptr.FirstName)
	fmt.Println(p.Age)
	ptr.FirstName = "lucas"
	increamentAge(ptr)
	fmt.Printf(p.FirstName)
	fmt.Println(p.Age)
}
