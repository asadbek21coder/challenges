package challenges

// Online Shopping System:
// Create a Customer struct with fields Name, Email, and Cart. Implement methods to add items to the cart, remove items from the cart, and display the cart contents.

import "fmt"

type Item struct {
	Name  string
	Price float64
}

type Customer struct {
	Name  string
	Email string
	Cart  []Item
}

func (c *Customer) AddToCart(item Item) {
	c.Cart = append(c.Cart, item)
	fmt.Println("Item added to cart.")
}

func (c *Customer) RemoveFromCart(itemName string) {
	var updatedCart []Item
	for _, item := range c.Cart {
		if item.Name != itemName {
			updatedCart = append(updatedCart, item)
		}
	}
	c.Cart = updatedCart
	fmt.Println("Item removed from cart.")
}

func (c Customer) DisplayCart() {
	fmt.Printf("Customer: %s\n", c.Name)
	fmt.Println("Cart contents:")
	for _, item := range c.Cart {
		fmt.Println("Item:", item.Name)
		fmt.Println("Price:", item.Price)
		fmt.Println()
	}
}

// func main() {
// 	customer := Customer{
// 		Name:  "John Doe",
// 		Email: "john.doe@example.com",
// 	}
// 	customer.AddToCart(Item{Name: "Item 1", Price: 9.99})
// 	customer.AddToCart(Item{Name: "Item 2", Price: 19.99})
// 	customer.AddToCart(Item{Name: "Item 3", Price: 14.99})

// 	customer.DisplayCart()

// 	customer.RemoveFromCart("Item 2")

// 	customer.DisplayCart()
// }
// Output:

// vbnet
// Copy code
// Item added to cart.
// Item added to cart.
// Item added to cart.
// Customer: John Doe
// Cart contents:
// Item: Item 1
// Price: 9.99

// Item: Item 2
// Price: 19.99

// Item: Item 3
// Price: 14.99

// Item removed from cart.
// Customer: John Doe
// Cart contents:
// Item: Item 1
// Price: 9.99

// Item: Item 3
// Price: 14.99
