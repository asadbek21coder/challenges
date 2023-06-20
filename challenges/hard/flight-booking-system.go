package challenges

// Flight Booking System:
// Create a Flight struct with fields FlightNumber, Source, Destination, and Passengers. Implement methods to add passengers, remove passengers, and display the passenger list.

import "fmt"

type Passenger struct {
	Name  string
	Age   int
	Seat  string
	Class string
}

type Flight struct {
	FlightNumber string
	Source       string
	Destination  string
	Passengers   []Passenger
}

func (f *Flight) AddPassenger(passenger Passenger) {
	f.Passengers = append(f.Passengers, passenger)
	fmt.Println("Passenger added.")
}

func (f *Flight) RemovePassenger(passengerName string) {
	var updatedPassengers []Passenger
	for _, passenger := range f.Passengers {
		if passenger.Name != passengerName {
			updatedPassengers = append(updatedPassengers, passenger)
		}
	}
	f.Passengers = updatedPassengers
	fmt.Println("Passenger removed.")
}

func (f Flight) DisplayPassengerList() {
	fmt.Println("Passenger list:")
	for _, passenger := range f.Passengers {
		fmt.Printf("Name: %s, Age: %d, Seat: %s, Class: %s\n", passenger.Name, passenger.Age, passenger.Seat, passenger.Class)
	}
}

// func main() {
// 	flight := Flight{
// 		FlightNumber: "F123",
// 		Source:       "New York",
// 		Destination:  "London",
// 	}

// 	passenger1 := Passenger{Name: "John Doe", Age: 30, Seat: "A1", Class: "Business"}
// 	passenger2 := Passenger{Name: "Jane Smith", Age: 25, Seat: "B2", Class: "Economy"}

// 	flight.AddPassenger(passenger1)
// 	flight.AddPassenger(passenger2)

// 	flight.DisplayPassengerList()

// 	flight.RemovePassenger("John Doe")

// 	flight.DisplayPassengerList()
// }
// Output:

// yaml
// Copy code
// Passenger added.
// Passenger added.
// Passenger list:
// Name: John Doe, Age: 30, Seat: A1, Class: Business
// Name: Jane Smith, Age: 25, Seat: B2, Class: Economy
// Passenger removed.
// Passenger list:
// Name: Jane Smith, Age: 25, Seat: B2, Class: Economy
