package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var count int64
	atomic.AddInt64(&count, 1)
	fmt.Println(count)
}
