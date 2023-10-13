package main

import "fmt"

// s := arr[start:end]

func main() {
	var colors = [5]string{
		"Red",
		"Blue",
		"Green",
		"Pink",
		"Purple",
	}
	s1 := colors[1:3] // Selected Indices : 1,2

	s2 := colors[:3] // Selected Indices : 0,1,2

	s3 := colors[2:] // Selected Indices : 2,3,4

	s4 := colors[:] // Selected Indices : 0,1,2,3,4

	fmt.Println("Array:", colors)
	fmt.Println("Slice 1:", s1)
	fmt.Println("Slice 2:", s2)
	fmt.Println("Slice 3:", s3)
	fmt.Println("Slice 4:", s4)
}
