package main

import "fmt"

type operationFunctionType = func(num1, num2 int) int

func printResult(f operationFunctionType) {
	res := f(10, 20)
	fmt.Println(res)
}

func main() {
	var sum operationFunctionType = func(num1, num2 int) int {
		return num1 + num2
	}
	var multiply operationFunctionType = func(num1, num2 int) int {
		return num1 * num2
	}

	var subtraction operationFunctionType = func(num1, num2 int) int {
		return num1 - num2
	}

	printResult(sum)
	printResult(multiply)
	printResult(subtraction)
}
