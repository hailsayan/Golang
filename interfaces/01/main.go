package main

import "fmt"

type Bosch struct{}

type GPlus struct{}

func (machine Bosch) Cleaning() {

	fmt.Println("Bosch Cleaning")

}

func (machine Bosch) Drying() {

	fmt.Println("Bosch Drying")

}

func (machine GPlus) Cleaning() {

	fmt.Println("GPlus Cleaning")

}

func (machine GPlus) Drying() {

	fmt.Println("GPlus Drying")

}

type WashingMachine interface {
	Cleaning()

	Drying()
}

func CleanAndDry(machine WashingMachine) {

	machine.Cleaning()

	machine.Drying()

}

func main() {
	b := Bosch{}
	gp := GPlus{}

	CleanAndDry(b)
	CleanAndDry(gp)
}
