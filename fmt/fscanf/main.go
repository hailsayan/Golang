package main

import (
	"fmt"
	"os"
)

func main() {
	var d int
	fmt.Fscanf(os.Stdin, "%d", &d)
	fmt.Println(d)
}
