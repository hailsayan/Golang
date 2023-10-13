package main

import "fmt"

func main() {
	metals := []string{"gold", "silver"}
	fmt.Printf("len : %d , cap : %d\n", len(metals), cap(metals))

	metals = append(metals, "iron")
	fmt.Printf("len : %d , cap : %d\n", len(metals), cap(metals))
}
