package main

import "fmt"

func changeVarByAddress(addr *float64) {
	*addr += 1.1
}

func main() {
	a := 4.4
	changeVarByAddress(&a)
	fmt.Println(a)
}
