package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("Goroutine:", i)
			time.Sleep(time.Millisecond * 300)
		}
	}()
	for i := 0; i < 3; i++ {
		fmt.Println("Main:", i)
		time.Sleep(time.Millisecond * 300)
	}
	time.Sleep(1 * time.Second)
}
