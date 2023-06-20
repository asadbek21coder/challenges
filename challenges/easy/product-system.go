package challenges

// Product2 Inventory System:
// Create a struct called Product2 with fields name, price, and quantity. Implement methods to update the quantity, calculate the total value, and display product information.

import "fmt"

type Product2 struct {
	name     string
	price    float64
	quantity int
}

func (p *Product2) UpdateQuantity(quantity int) {
	p.quantity += quantity
}

func (p Product2) CalculateTotalValue() float64 {
	return p.price * float64(p.quantity)
}

func (p Product2) DisplayInformation() {
	fmt.Println("Name:", p.name)
	fmt.Println("Price:", p.price)
	fmt.Println("Quantity:", p.quantity)
	fmt.Println("Total Value:", p.CalculateTotalValue())
}

// func main() {
// 	product := Product2{
// 		name:     "iPhone",
// 		price:    999.99,
// 		quantity: 5,
// 	}
// 	product.UpdateQuantity(2)
// 	product.DisplayInformation()
// }
// Output:

// yaml
// Copy code
// Name: iPhone
// Price: 999.99
// Quantity: 7
// Total Value: 6999.93
