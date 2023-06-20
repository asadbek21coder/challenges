package challenges

// Rectangle Geometry:
// Create a Rectangle struct with fields length and width. Implement methods to calculate the area and perimeter of the rectangle.

type Rectangle struct {
	length float64
	width  float64
}

func (r Rectangle) CalculateArea() float64 {
	return r.length * r.width
}

func (r Rectangle) CalculatePerimeter() float64 {
	return 2 * (r.length + r.width)
}

// func main() {
// 	rectangle := Rectangle{length: 5.0, width: 3.0}
// 	area := rectangle.CalculateArea()
// 	perimeter := rectangle.CalculatePerimeter()

// 	fmt.Println("Area:", area)
// 	fmt.Println("Perimeter:", perimeter)
// }

// Output:

// makefile
// Copy code
// Area: 15
// Perimeter: 16
