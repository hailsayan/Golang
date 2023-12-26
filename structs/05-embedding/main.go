package main

import "fmt"

type Product struct {
	Name  string
	Price int
}
type Electrical struct {
	Product
	WarrantyMonths int
}
type Clothing struct {
	Product
	Size     string
	Material string
}

func main() {
	cl := Clothing{}
	cl.Name = "shirt"
	cl.Price = 300
	cl.Size = "L"
	cl.Material = "cotton"
	fmt.Printf("%+v\n", cl)

	el := Electrical{Product: Product{Name: "lamp", Price: 40}, WarrantyMonths: 24}
	fmt.Printf("%+v\n", el)
}
