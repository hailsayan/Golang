package main

import (
	"fmt"
	"time"
)

func sendData(ch chan<- string, txt string) {
	ch <- txt
}
func receiveData(ch <-chan string) {
	fmt.Println(<-ch)
}
func main() {
	ch := make(chan string, 1)
	go sendData(ch, "Message")
	go receiveData(ch)
	time.Sleep(1 * time.Second)
}
