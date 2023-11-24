package main

import "fmt"

func main() {
	printName := func(name string) string {
		return fmt.Sprintf("your name : %s", name)
	}

	fmt.Println(printname("Psyon"))
}
