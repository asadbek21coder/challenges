package challenges

// Online Store:
// Create a Product struct with fields Name, Price, and Quantity. Implement methods to display the product details, add stock, and sell products.

import "fmt"

type Product struct {
	Name     string
	Price    float64
	Quantity int
}

func (p Product) DisplayDetails() {
	fmt.Println("Name:", p.Name)
	fmt.Println("Price:", p.Price)
	fmt.Println("Quantity:", p.Quantity)
}

func (p *Product) AddStock(quantity int) {
	p.Quantity += quantity
	fmt.Println("Stock added successfully.")
}

func (p *Product) Sell(quantity int) {
	if quantity > p.Quantity {
		fmt.Println("Insufficient stock.")
	} else {
		p.Quantity -= quantity
		fmt.Println("Sale successful.")
	}
}
