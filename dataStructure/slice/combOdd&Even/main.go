package main

import "fmt"

// ... is used to unpack a slice and pass its individual elements to a variadic function or append its elements to another slice.
func main() {
	even := []int{2, 4, 6}
	odd := []int{1, 3, 5}
	numbers := append(odd, even...)
	fmt.Println(numbers)
}
