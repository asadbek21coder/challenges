package challenges

import "math"

// Shape:
// Create a struct called Shape with a method to calculate the area. Implement derived structs Rectangle2 and Circle that inherit from Shape and override the CalculateArea() method.

type Shape interface {
	CalculateArea() float64
}

type Rectangle2 struct {
	width  float64
	height float64
}

func (r Rectangle2) CalculateArea() float64 {
	return r.width * r.height
}

type Circle struct {
	radius float64
}

func (c Circle) CalculateArea() float64 {
	return math.Pi * c.radius * c.radius
}

// func main() {
// 	rectangle := Rectangle2{width: 5.0, height: 3.0}
// 	circle := Circle{radius: 2.5}

// 	shapes := []Shape{rectangle, circle}

// 	for _, shape := range shapes {
// 		area := shape.CalculateArea()
// 		fmt.Printf("Area: %.2f\n", area)
// 	}
// }
// Output:

// makefile
// Copy code
// Area: 15.00
// Area: 19.63
