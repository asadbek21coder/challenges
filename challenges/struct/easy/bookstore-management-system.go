package challenges

import "fmt"

// Bookstore Management System:

// Create a `Book` struct with fields `Title`, `Author`, and `Price`.
//  Implement methods to apply a discount to the book price and display book details.

type Book struct {
	Title  string
	Author string
	Price  float64
}

func (b *Book) ApplyDiscount(discountPercentage float64) {
	b.Price = b.Price - b.Price*(discountPercentage/100)
}

func (b Book) Display() {
	fmt.Println("Title:", b.Title)
	fmt.Println("Author:", b.Author)
	fmt.Println("Price:", b.Price)
}

// func main() {

// }
