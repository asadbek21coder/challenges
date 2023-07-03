package challenges

import "fmt"

// Product2 Inventory System:

// Create a struct called Product2 with fields name,
// price, and quantity. Implement methods to update the
// quantity, calculate the total value, and display product
// information.

// Mahsulot - nomi, narxi, miqdori
// Miqdorini o'zgartirish - Kopaytirish(), Kamaytirish()
// Jami() summani hisoblash

type Mahsulot struct {
	Nomi    string
	Narxi   float64
	Miqdori int
}

func (m *Mahsulot) Kopaytirish(dona int) {
	m.Miqdori += dona
}

func (m *Mahsulot) Kamaytirish(dona int) {
	if m.Miqdori >= dona {
		m.Miqdori -= dona
	}
}

func (m Mahsulot) Jami() float64 {
	return float64(m.Miqdori) * m.Narxi
}

func (m Mahsulot) Display() {
	fmt.Println("Mahsulot nomi: ", m.Nomi)
	fmt.Println("Mahsulot narxi: ", m.Narxi)
	fmt.Println("Mahsulot miqdori: ", m.Miqdori)
	fmt.Println("Jami: ", m.Jami(), "so'mlik mahsulot bor!")
}
