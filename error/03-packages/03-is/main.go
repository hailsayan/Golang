package main

import (
	"errors"
	"fmt"
)

var ErrBadInput = errors.New("bad input")

func validateInput(input string) error {
	if len(input) == 0 {
		return errors.Join(fmt.Errorf("Invalid Error 2"), errors.New("Invalid Error 1"), ErrBadInput)
	}
	return nil
}
func main() {
	err := validateInput("")
	fmt.Println(errors.Is(err, ErrBadInput))
}
