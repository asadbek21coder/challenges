package hard

import (
	"fmt"
	"time"
)

// Bank Account System:

// Create a BankAccount2 struct with fields AccountNumber,
// HolderName, Balance, and Transactions.
// Implement methods to deposit money, withdraw money,
// and display the transaction history.

// Transaction -
// 		Amount     int
//      Transacted time.Time
// BankAccount2 struct - AccountNumber,
// HolderName, Balance, and Transactions array

type Transaction struct {
	Amount     float64
	Transacted time.Time
}

type BankAccount2 struct {
	AccountNumber string
	HolderName    string
	Balance       float64
	Transactions  []Transaction
}

func (b *BankAccount2) Deposit(income float64) {
	b.Balance = b.Balance + income
	newTransaction := Transaction{
		Amount:     income,
		Transacted: time.Now(),
	}
	fmt.Println("Successfully transacted!")

	b.Transactions = append(b.Transactions, newTransaction)
}

func (b *BankAccount2) Withdraw(outcome float64) {
	if outcome < b.Balance {
		b.Balance = b.Balance - outcome
		fmt.Println("Successfully transacted!")

		newTransaction := Transaction{
			Amount:     -outcome,
			Transacted: time.Now(),
		}
		b.Transactions = append(b.Transactions, newTransaction)
	} else {
		fmt.Println("Insufficient money!")
	}
}

func (b BankAccount2) Display() {
	var positive, negative []Transaction
	for _, v := range b.Transactions {
		if v.Amount < 0 {
			negative = append(negative, v)
		} else {
			positive = append(positive, v)
		}
	}

	fmt.Println("Hosib raqami: ", b.AccountNumber)
	fmt.Println("Egasi: ", b.HolderName)
	fmt.Println("Miqdori: ", b.Balance)
	fmt.Println("O'tkazmalar: ")
	fmt.Println(positive, negative)

}
