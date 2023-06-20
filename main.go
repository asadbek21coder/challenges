package main

import challenges "github.com/asadbek21coder/challenges/challenges/easy"

func main() {
	account := challenges.BankAccount{
		AccountNumber: "123456789",
		HolderName:    "John Doe",
		Balance:       1000.0,
	}

	account.Deposit(500.0)
	account.Withdraw(200.0)

	account.Display()
}
