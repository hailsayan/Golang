package main

import "fmt"

func applyFunction(f func(int) int, num int) int {
	return f(num)
}
func square(n int) int {
	return n * n
}
func cube(n int) int {
	return n * n * n
}
func main() {
	fmt.Println(applyFunction(square, 5))
	fmt.Println(applyFunction(cube, 5))
}
