package main

import (
	"errors"
	"fmt"
)

type MyError struct {
	Code    int
	Message string
}

func (e MyError) Error() string {
	return fmt.Sprint("MyError Error")
}

type person struct {
	Name string
	Age  int
}

func (p person) Error() string {
	return fmt.Sprint("Person Error")
}
func main() {
	rootErr := MyError{Code: 404, Message: "Not Found"}
	var otherErr *person
	fmt.Println(errors.As(rootErr, &otherErr))
}
