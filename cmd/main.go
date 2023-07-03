package main

import (
	"fmt"

	"github.com/asadbek21coder/algo/challenges/mix"
)

func main() {
	fmt.Println(mix.OnlyVowels("kons+.qpirolog*/-=+oya"))
	// fmt.Println(mix.NumberOfVovels("maarifatparvar"))
	// fmt.Println(mix.SaleKonfet(16, 4))
	// distance, err := mix.DistanceBoat(6, 7, 2, 4)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(distance)
	// }

	// fmt.Println(mix.Distance(10, 35, 3))
	// fmt.Println(mix.Difference(8.5, 5, 6000, 16000))
	// var calculator algos.Calculator

	// calculator = algos.Rectangle{
	// 	Height: 12,
	// 	Width:  13,
	// }

	// fmt.Printf("%T, %v\n", calculator, calculator)
	// fmt.Println(calculator.Area(), calculator.Perimetre())

	// calculator = algos.Triangle{
	// 	A: 3,
	// 	B: 4,
	// 	C: 5,
	// }
	// fmt.Printf("%T, %v\n", calculator, calculator)
	// fmt.Println(calculator.Area(), calculator.Perimetre())

	// lbr := challenges.Library{
	// 	Name:  "Alisher Navoiy",
	// 	Books: []challenges.Book{},
	// }

	// kitob1 := challenges.Book{
	// 	Name:     "Harry Potter and Order of Phenix",
	// 	Author:   "J.K.Rowling",
	// 	Quantity: 2,
	// }
	// kitob2 := challenges.Book{
	// 	Name:     "Atomic Habits",
	// 	Author:   "James Clear",
	// 	Quantity: 3,
	// }
	// kitob3 := challenges.Book{
	// 	Name:     "Eljernonga atalgan gullar",
	// 	Author:   "Daniel Kiz",
	// 	Quantity: 0,
	// }

	// lbr.AddBook(kitob1)
	// lbr.AddBook(kitob1)
	// lbr.AddBook(kitob2)
	// lbr.AddBook(kitob2)
	// lbr.AddBook(kitob2)
	// lbr.AddBook(kitob3)
	// lbr.AddBook(kitob3)
	// lbr.LendBook(kitob1)
	// lbr.LendBook(kitob1)
	// lbr.LendBook(kitob1)
	// lbr.LendBook(kitob1)
	// lbr.LendBook(kitob1)

	// lbr.Display()
	// mahsulot1 := challenges.Mahsulot{
	// 	Nomi:    "Olma",
	// 	Narxi:   14,
	// 	Miqdori: 57,
	// }

	// mahsulot1.Kopaytirish(20)
	// mahsulot1.Display()
	// mahsulot1.Kamaytirish(5)
	// mahsulot1.Display()

	// em1 := challenges.Employee{
	// 	Name:   "Muhammad",
	// 	Age:    22,
	// 	Salary: 1200.0,
	// }

	// em1.Display()
	// em1.IncreaseSalary(20)
	// em1.Display()

	// ticket1 := challenges.Ticket{
	// 	TicketNumber: 1,
	// 	EventName:    "Conference - One belt, one way",
	// 	Price:        10.8,
	// 	Status:       "Available",
	// }

	// ticket1.Display()
	// ticket1.BookTicket()
	// ticket1.Display()
	// ticket1.CancelTicket()
	// ticket1.Display()
	// ticket1.BookTicket()
	// ticket1.Display()

	// account := challenges.BankAccount{
	// 	AccountNumber: "123456789",
	// 	HolderName:    "John Doe",
	// 	Balance:       1000.0,
	// }

	// account.Deposit(500.0)
	// account.Withdraw(200.0)

	// account.Display()

	// book := challenges.Book{
	// 	Title:  "The Great Gatsby",
	// 	Author: "F. Scott Fitzgerald",
	// 	Price:  300,
	// }

	// book.ApplyDiscount(10.0)
	// book.ApplyDiscount(10.0)

	// book.Display()

	// car := challenges.Car{
	// 	Make:       "Toyota",
	// 	Model:      "Corolla",
	// 	RentPerDay: 40.0,
	// }

	// fmt.Println("10 kunga: ", car.CalculateRent(14), "$")
	// car.Display()

	// mijoz1 := challenges.Mijoz{
	// 	Ismi:     "Muhammad",
	// 	Email:    "muhammad@gmail.com",
	// 	Savatcha: []challenges.MahsulotTuri{},
	// }
	// mijoz2 := challenges.Mijoz{
	// 	Ismi:     "Bahrom",
	// 	Email:    "bahrom@gmail.com",
	// 	Savatcha: []challenges.MahsulotTuri{},
	// }

	// sut := challenges.MahsulotTuri{
	// 	Nomi: "Sut",
	// 	Narx: 15000.0,
	// }
	// non := challenges.MahsulotTuri{
	// 	Nomi: "Non",
	// 	Narx: 3000.0,
	// }
	// chips := challenges.MahsulotTuri{
	// 	Nomi: "Chips",
	// 	Narx: 10000.0,
	// }

	// mijoz1.SavatchagaQosh(sut)
	// mijoz1.SavatchagaQosh(non)
	// mijoz1.SavatchagaQosh(chips)
	// // fmt.Println(mijoz1)
	// mijoz1.SavatchadanOchir(chips)
	// fmt.Println(mijoz1)
	// mijoz2.SavatchagaQosh(sut)
	// mijoz2.SavatchagaQosh(non)
	// mijoz2.SavatchagaQosh(chips)
	// fmt.Println(mijoz2)

	// asad := challenges.User{
	// 	Name:    "Asadbek",
	// 	Email:   "asadbek@gmail.com",
	// 	Friends: []challenges.User{},
	// }

	// muhammad := challenges.User{
	// 	Name:    "Muhammad",
	// 	Email:   "muhammad@gmail.com",
	// 	Friends: []challenges.User{},
	// }

	// saidkamol := challenges.User{
	// 	Name:    "Saidkamol",
	// 	Email:   "saidkamol@gmail.com",
	// 	Friends: []challenges.User{},
	// }
	// nodirbek := challenges.User{
	// 	Name:    "Nodirbek",
	// 	Email:   "nodirbek@gmail.com",
	// 	Friends: []challenges.User{},
	// }

	// asad.AddFriends(nodirbek)
	// asad.AddFriends(saidkamol)
	// saidkamol.AddFriends(asad)
	// saidkamol.AddFriends(muhammad)
	// muhammad.AddFriends(saidkamol)
	// muhammad.AddFriends(asad)
	// muhammad.AddFriends(nodirbek)
	// nodirbek.AddFriends(saidkamol)

	// // muhammad.Display()
	// var hozir time.Time
	// fmt.Println(hozir)
	// fmt.Printf("%T\n", hozir)
	// hozir = time.Now()
	// // formatted := hozir.Format("RFC1123Z")
	// fmt.Println(hozir.Format(time.RFC1123))

	// bankAccount := hard.BankAccount2{
	// 	AccountNumber: "0001",
	// 	HolderName:    "Bahrom",
	// 	Balance:       1000,
	// 	Transactions:  []hard.Transaction{},
	// }

	// bankAccount.Deposit(200)
	// bankAccount.Withdraw(300)
	// bankAccount.Display()

	// fmt.Println(bankAccount)

	// reys1 := hard.Reys{
	// 	ReysRaqami:  001,
	// 	Qayerdan:    "Toshkent",
	// 	Qayerga:     "Urganch",
	// 	Yolovchilar: []hard.Client{},
	// }

	// muhammad := hard.Client{
	// 	Ismi:  "Muhammad",
	// 	Yoshi: 21,
	// 	Orni:  10,
	// }
	// bahrom := hard.Client{
	// 	Ismi:  "Bahrom",
	// 	Yoshi: 23,
	// 	Orni:  11,
	// }
	// bahodir := hard.Client{
	// 	Ismi:  "Bahodir",
	// 	Yoshi: 21,
	// 	Orni:  12,
	// }

	// Izzat := hard.Client{
	// 	Ismi:  "Izzat",
	// 	Yoshi: 21,
	// 	Orni:  13,
	// }
	// saidkamol := hard.Client{
	// 	Ismi:  "Saidkamol",
	// 	Yoshi: 19,
	// 	Orni:  18,
	// }
	// reys1.YolovchiQoshish(muhammad)
	// reys1.YolovchiQoshish(bahrom)
	// reys1.YolovchiQoshish(Izzat)
	// reys1.YolovchiQoshish(bahodir)
	// reys1.YolovchiQoshish(saidkamol)

	// // fmt.Println(reys1)
	// reys1.Display()

	// student1 := challenges.Student{
	// 	Name:   "Muhammad",
	// 	Age:    21,
	// 	Grades: []int{4, 3, 5},
	// }

	// student1.AddGrade(4)
	// student1.AddGrade(5)
	// student1.AddGrade(3)
	// student1.AddGrade(2)

	// student1.Display()

	// fmt.Printf("Avarage grade of the student: %.1f\n", student1.CalculateAvarageGrade())
}

// func Area(a, b float64) float64 {
// 	// area := a * b
// 	// return area
// 	return a * b
// }

// func Perimetre(a, b float64) float64 {
// 	// p := a + a + b + b
// 	p := 2*a + 2*b
// 	// p := 2 * (a + b)
// 	return p
// }
