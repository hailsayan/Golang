package main

import "fmt"

func main() {
	type Student struct {
		firstName string
		lastName  string
		age       int
	}

	s := Student{
		firstName: "john",
		lastName:  "doe",
		age:       25,
	}

	fmt.Printf("lastName : %s\n", s.lastName)
	fmt.Printf("age : %d", s.age)
}
