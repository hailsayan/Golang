package main

import "fmt"

func introduce(name string, age int) {
	fmt.Printf("My Name is %s and I'm %d years old", name, age)
}

func main() {
	introduce("psyon", 20)
}
