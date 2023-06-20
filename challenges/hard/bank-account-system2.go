package challenges

// Bank Account System:
// Create a BankAccount2 struct with fields AccountNumber, HolderName, Balance, and Transactions. Implement methods to deposit money, withdraw money, and display the transaction history.

import (
	"fmt"
	"time"
)

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

func (ba *BankAccount2) Deposit(amount float64) {
	ba.Balance += amount
	ba.Transactions = append(ba.Transactions, Transaction{Amount: amount, Transacted: time.Now()})
	fmt.Println("Deposit successful.")
}

func (ba *BankAccount2) Withdraw(amount float64) {
	if amount <= ba.Balance {
		ba.Balance -= amount
		ba.Transactions = append(ba.Transactions, Transaction{Amount: -amount, Transacted: time.Now()})
		fmt.Println("Withdrawal successful.")
	} else {
		fmt.Println("Insufficient balance.")
	}
}

func (ba BankAccount2) DisplayTransactionHistory() {
	fmt.Printf("Transaction history for Account %s:\n", ba.AccountNumber)
	for _, transaction := range ba.Transactions {
		fmt.Printf("Amount: %.2f, Transacted: %s\n", transaction.Amount, transaction.Transacted.Format("2006-01-02 15:04:05"))
	}
}

// func main() {
// 	account := BankAccount2{
// 		AccountNumber: "123456789",
// 		HolderName:    "John Doe",
// 		Balance:       1000.0,
// 	}

// 	account.Deposit(500.0)
// 	account.Withdraw(200.0)

// 	account.DisplayTransactionHistory()
// }
