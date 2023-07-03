package challenges

import "fmt"

// Create a Student struct with fields Name, Age, and Grade.
// Implement methods to promote the student to the next grade
// and display student details.

// Student - Name, Age, Grades[]

// AddGrade()
// CalculateAvarageGrade()
// Display()

// Talaba - Ismi, Yoshi, Baholari[]
// BahoQoshish()
// OrtachaBahoniHisoblash
// Korsatish()

type Student struct {
	Name   string
	Age    int
	Grades []int
}

func (s *Student) AddGrade(grade int) {
	s.Grades = append(s.Grades, grade)
}

func (s Student) CalculateAvarageGrade() float64 {
	var sum float64
	for _, v := range s.Grades {
		sum += float64(v)
	}
	return sum / float64(len(s.Grades))
}

func (s Student) Display() {
	fmt.Println("\nName: ", s.Name)
	fmt.Println("Age: ", s.Age)
	fmt.Println("Grades: ")
	for _, v := range s.Grades {
		fmt.Println("\t", v)
	}
}
