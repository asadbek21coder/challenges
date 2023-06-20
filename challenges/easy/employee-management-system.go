package challenges

// Employee Management System:
// Create an Employee struct with fields Name, Age, and Salary. Implement methods to increase salary and display employee details.
import "fmt"

type Employee struct {
	Name   string
	Age    int
	Salary float64
}

func (e *Employee) IncreaseSalary(percentage float64) {
	e.Salary += e.Salary * (percentage / 100)
}

func (e Employee) Display() {
	fmt.Println("Name:", e.Name)
	fmt.Println("Age:", e.Age)
	fmt.Println("Salary:", e.Salary)
}

// func main() {
// 	employee := Employee{
// 		Name:   "John Doe",
// 		Age:    30,
// 		Salary: 5000.0,
// 	}

// 	employee.IncreaseSalary(10.0)

// 	employee.Display()
// }
