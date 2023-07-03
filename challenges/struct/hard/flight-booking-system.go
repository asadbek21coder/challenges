package hard

import "fmt"

// Flight Booking System:

// Create a Flight struct with fields FlightNumber,
// Source, Destination, and Passengers.
// Implement methods
// to add passengers, remove passengers, and display the
// passenger list.

// Parvoz - ParvozRaqami, Qayerdan,Qayerga, Yolovchilar[]
// Yolovchi - Ismi, Yoshi, Orni
// YolovchiQoshish() - tekshiriladi, o'rin band emasligi
// YolovchiBekorQilish()
// YolovchilarListiniKorsatish()

type Client struct {
	Ismi  string
	Yoshi int
	Orni  int
}

type Reys struct {
	ReysRaqami  int
	Qayerdan    string
	Qayerga     string
	Yolovchilar []Client
}

func (r *Reys) YolovchiQoshish(yolovchi Client) {
	for _, v := range r.Yolovchilar {
		if v.Orni == yolovchi.Orni {
			fmt.Println("Bu o'rin band")
			return
		}
	}
	r.Yolovchilar = append(r.Yolovchilar, yolovchi)
}

func (r Reys) Display() {

	fmt.Println("Reys raqami: ", r.ReysRaqami)
	fmt.Println("Qayerdan: ", r.Qayerdan)
	fmt.Println("Qayerga: ", r.Qayerga)
	fmt.Println("Yo'ovchilar: ")
	for _, v := range r.Yolovchilar {
		fmt.Println("\t Ismi: ", v.Ismi)
		fmt.Println("\t Yoshi: ", v.Yoshi)
		fmt.Println("\t O`rni: ", v.Orni)
		fmt.Println("______________________")
	}
}
