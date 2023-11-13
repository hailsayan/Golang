package main

import "fmt"

func swap(num1, num2 int) (int, int) {
	return num2, num1
}

func main() {
	num1 := 5
	num2 := 20

	num1, num2 = swap(num1, num2)
	fmt.Printf("num1 : %d, num2 : %d", num1, num2)
}
