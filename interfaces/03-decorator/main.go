package main

type Pizza interface {
	getPrice() int
}

type Veggie struct{}

func (p *Veggie) getPrice() int {

	return 15

}
