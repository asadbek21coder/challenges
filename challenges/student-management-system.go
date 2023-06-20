package challenges

// Create a Student struct with fields Name, Age, and Grade. Implement methods to promote the student to the next grade and display student details.

import "fmt"

type Student struct {
	Name  string
	Age   int
	Grade int
}

func (s *Student) Promote() {
	s.Grade++
}

func (s Student) Display() {
	fmt.Println("Name:", s.Name)
	fmt.Println("Age:", s.Age)
	fmt.Println("Grade:", s.Grade)
}

// func main() {
// 	student := Student{
// 		Name:  "John Doe",
// 		Age:   15,
// 		Grade: 9,
// 	}

// 	student.Promote()

// 	student.Display()
// }
