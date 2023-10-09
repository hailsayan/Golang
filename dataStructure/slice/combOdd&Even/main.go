package main

import "fmt"

func main() {
	even := []int{2, 4, 6}
	odd := []int{1, 3, 5}
	numbers := append(odd, even...) // ...

	fmt.Println(numbers)
}
