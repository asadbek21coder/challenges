package challenges

import "fmt"

// Online Ticket Booking System:
// Create a Ticket struct with fields TicketNumber,
// EventName, Price, and Status. Implement methods to book a
// ticket, cancel a ticket, and display ticket details.

type Ticket struct {
	TicketNumber int
	EventName    string
	Price        float64
	Status       string
}

func (t *Ticket) BookTicket() {
	if t.Status == "Available" || t.Status == "Canceled" {
		t.Status = "Booked"
	}
}

func (t *Ticket) CancelTicket() {
	if t.Status == "Booked" {
		t.Status = "Canceled"
	}
}

func (t Ticket) Display() {
	fmt.Println(t.TicketNumber)
	fmt.Println(t.EventName)
	fmt.Println(t.Price)
	fmt.Println(t.Status)
}
