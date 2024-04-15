package main

import (
	"errors"
	"fmt"
)

type BadInputError struct {
	input string
}

func (e *BadInputError) Error() string {
	return fmt.Sprintf("bad input is: %s", e.input)
}
func validateInput(input string) error {
	if len(input) == 0 {
		return errors.Join(fmt.Errorf("Bad Input Error One"), fmt.Errorf("Bad Input Error Two"), &BadInputError{input: input})
	}
	return nil
}
func main() {
	err := validateInput("")
	var badInputErr *BadInputError
	fmt.Println(errors.As(err, &badInputErr))
}
