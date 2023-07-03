package challenges

import "fmt"

// Employee Management System:
// Create an Employee struct with fields Name, Age, and Salary.
//  Implement methods to increase salary and
// display employee details.

type Employee struct {
	Name   string
	Age    int
	Salary float64
}

func (e *Employee) IncreaseSalary(procent float64) {
	e.Salary = e.Salary + procent*e.Salary/100
}

func (e Employee) Display() {
	fmt.Println("Name: ", e.Name)
	fmt.Println("Age: ", e.Age)
	fmt.Println("Salary: ", e.Salary, "$")
}
