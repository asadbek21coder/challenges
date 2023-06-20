package challenges

// Student2 Gradebook:
// Create a Student2 struct with fields name and grades. Implement a method to display the student's name and average grade.

import "fmt"

type Student2 struct {
	name   string
	grades []int
}

func (s Student2) DisplayDetails() {
	fmt.Println("Name:", s.name)
	fmt.Println("Average Grade:", s.calculateAverageGrade())
}

func (s Student2) calculateAverageGrade() float64 {
	sum := 0
	for _, grade := range s.grades {
		sum += grade
	}
	return float64(sum) / float64(len(s.grades))
}

// func main() {
// 	student := Student2{
// 		name:   "John Doe",
// 		grades: []int{80, 85, 90, 95},
// 	}
// 	student.DisplayDetails()
// }
// Output:

// yaml
// Copy code
// Name: John Doe
// Average Grade: 87.5
