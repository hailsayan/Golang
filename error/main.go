package main

import (
	"errors"
	"fmt"
)

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("error : division by zero")
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
