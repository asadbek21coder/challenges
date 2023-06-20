package challenges

// Online Ticket Booking System:
// Create a Ticket struct with fields TicketNumber, EventName, Price, and Status. Implement methods to book a ticket, cancel a ticket, and display ticket details.

import "fmt"

type Ticket struct {
	TicketNumber int
	EventName    string
	Price        float64
	Status       string
}

func (t *Ticket) BookTicket() {
	if t.Status == "Available" {
		t.Status = "Booked"
		fmt.Println("Ticket booked successfully.")
	} else {
		fmt.Println("Ticket not available for booking.")
	}
}

func (t *Ticket) CancelTicket() {
	if t.Status == "Booked" {
		t.Status = "Cancelled"
		fmt.Println("Ticket cancelled successfully.")
	} else {
		fmt.Println("Ticket cannot be cancelled.")
	}
}

func (t Ticket) DisplayDetails() {
	fmt.Println("Ticket Number:", t.TicketNumber)
	fmt.Println("Event Name:", t.EventName)
	fmt.Println("Price:", t.Price)
	fmt.Println("Status:", t.Status)
}

// func main() {
// 	ticket := Ticket{
// 		TicketNumber: 1234,
// 		EventName:    "Concert",
// 		Price:        50.0,
// 		Status:       "Available",
// 	}

// 	ticket.BookTicket()
// 	ticket.DisplayDetails()

// 	ticket.CancelTicket()
// 	ticket.DisplayDetails()
// }
// Output:

// yaml
// Copy code
// Ticket booked successfully.
// Ticket Number: 1234
// Event Name: Concert
// Price: 50
// Status: Booked
// Ticket cancelled successfully.
// Ticket Number: 1234
// Event Name: Concert
// Price: 50
// Status: Cancelled
