package main

import (
	"fmt"
)

type DivisionError struct {
	StatusCode int
	Message    string
}

func (d DivisionError) Error() string {
	return fmt.Sprintf("code %d: %s", d.StatusCode, d.Message)
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, DivisionError{
			StatusCode: 2000,
			Message:    "cannot divide by zero",
		}
	}
	return a / b, nil
}
