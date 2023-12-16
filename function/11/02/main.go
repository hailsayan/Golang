package main

import "fmt"

func getMathFunction(operation string) func(int, int) int {
	switch operation {
	case "multiply":
		return func(a, b int) int {
			return a * b
		}
	default:
		return func(a, b int) int {
			return a + b
		}
	}
}

func main() {
	multiply := getMathFunction("multiply")
	fmt.Println(multiply(5, 25))
}
