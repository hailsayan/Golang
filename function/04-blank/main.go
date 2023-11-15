package main

import "fmt"

func returnTwoString() (string, string) {
	return "Hello", "World"
}

func main() {
	hello, - := returnTwoString()
	fmt.Println(hello)
}
