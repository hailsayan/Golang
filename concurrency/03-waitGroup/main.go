package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func mygoroutine() {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	fmt.Println("hello world")
}

func main() {
	wg.Add(1)
	go mygoroutine()
	wg.Wait()
	fmt.Println("finished")
}
