package main

// func panic(v interface{})

import (
	"fmt"
)

func main() {
	arr := make([]int, 0)
	arr[0] = 10
	fmt.Println(arr[4])
}
