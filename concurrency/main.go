package main

import (
	"fmt"
	"time"
)

func sayHello(name string) {
	fmt.Println("Hello " + name)
}
func main() {
	people := []string{"Rose", "Erich", "Amelia"}
	for _, person := range people {
		go sayHello(person)
	}
	time.Sleep(1 * time.Second)
}
