package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero = errors.New("division by zero")

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}
func main() {
	res, err := Divide(6, 0)
	if err != nil {
		fmt.Println(err)
		// Do something with the error
		return
	}
	fmt.Println(res)
}
