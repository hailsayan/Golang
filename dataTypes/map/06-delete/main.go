package main

import "fmt"

// delete (m,"hello")

var m = map[string]int{
	"hello": 0,
	"world": 1,
}

func main() {
	delete(m, "hello")
	fmt.Println(m)
}
