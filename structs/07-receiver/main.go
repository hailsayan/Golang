package main

import "fmt"

type Cat struct{}

func (c Cat) makeSound() {
	fmt.Println("meow")
}
func main() {
	pishi := Cat{}
	pishi.makeSound()
}
