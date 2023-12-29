package main

import "fmt"

type Cat struct {
	Color string
}

func (c Cat) ChangeColorByValue(color string) {

	c.Color = color

}

func (c *Cat) ChangeColorByReference(color string) {

	c.Color = color

}

func main() {
	snowy := Cat{Color: "white"}
	fmt.Println(snowy.Color)

	snowy.ChangeColorByValue("black")
	fmt.Println(snowy.Color)

	snowy.ChangeColorByReference("black")
	fmt.Println(snowy.Color)
}
