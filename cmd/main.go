package main

import (
	"github.com/asadbek21coder/challenges/challenges"
)

func main() {
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

	asad := challenges.User{
		Name:    "Asadbek",
		Email:   "asadbek@gmail.com",
		Friends: []challenges.User{},
	}

	muhammad := challenges.User{
		Name:    "Muhammad",
		Email:   "muhammad@gmail.com",
		Friends: []challenges.User{},
	}

	saidkamol := challenges.User{
		Name:    "Saidkamol",
		Email:   "saidkamol@gmail.com",
		Friends: []challenges.User{},
	}
	nodirbek := challenges.User{
		Name:    "Nodirbek",
		Email:   "nodirbek@gmail.com",
		Friends: []challenges.User{},
	}

	asad.AddFriends(nodirbek)
	asad.AddFriends(saidkamol)
	saidkamol.AddFriends(asad)
	saidkamol.AddFriends(muhammad)
	muhammad.AddFriends(saidkamol)
	muhammad.AddFriends(asad)
	muhammad.AddFriends(nodirbek)
	nodirbek.AddFriends(saidkamol)

	muhammad.Display()

}
