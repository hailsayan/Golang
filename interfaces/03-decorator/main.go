package main

import "fmt"

type Pizza interface {
	getPrice() int
}

type Veggie struct{}

func (p *Veggie) getPrice() int {

	return 15

}

type CheeseTopping struct {
	pizza Pizza
}

func (c *CheeseTopping) getPrice() int {

	pizzaPrice := c.pizza.getPrice()

	return pizzaPrice + 10

}

type TomatoTopping struct {
	pizza Pizza
}

func (c *TomatoTopping) getPrice() int {

	pizzaPrice := c.pizza.getPrice()

	return pizzaPrice + 5

}

func main() {

	pizza := &Veggie{}

	pizzaWithCheese := &CheeseTopping{

		pizza: pizza,
	}

	pizzaWithCheeseAndTomato := &TomatoTopping{

		pizza: pizzaWithCheese,
	}

	fmt.Printf("Price of Veggie Pizza with tomato and cheese topping is %d\n", pizzaWithCheeseAndTomato.getPrice())

}
