package main

import "fmt"

type Student struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	var s1 Student
	fmt.Printf("%+v", s1)
}
