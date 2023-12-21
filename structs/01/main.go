package main

import "fmt"

type Student struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	var s1 Student
	fmt.Printf("%+v\n", s1)

	var s2 Student = Student{
		firstName: "John",
		lastName:  "Doe",
		age:       25,
	}
	fmt.Printf("%+v\n", s2)

	var s3 = Student{
		firstName: "Psyon",
		age:       20,
	}
	fmt.Printf("%+v\n", s3)
}
