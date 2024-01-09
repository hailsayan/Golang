package main

import "fmt"

func main() {
	c := make(chan int, 2)
	go func() {
		c <- 20
	}()
	c <- 22
	fmt.Println(<-c)
	second := <-c
	fmt.Println(second)
}
