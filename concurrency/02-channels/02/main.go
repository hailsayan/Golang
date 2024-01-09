package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 10
	ch <- 20
	close(ch)

	for {
		value, ok := <-ch
		if !ok {
			fmt.Println("Channel closed")
			break
		} else {
			fmt.Println(value)
		}
	}
}
