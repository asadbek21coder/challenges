package challenges

// Bank Account System:
// Create a BankAccount struct with fields AccountNumber,
// HolderName, and Balance.
// Implement methods to deposit money, withdraw money,
// and display the account details.

import "fmt"

type BankAccount struct {
	AccountNumber string
	HolderName    string
	Balance       float64
}

func (ba *BankAccount) Deposit(amount float64) {
	ba.Balance += amount
}

func (ba *BankAccount) Withdraw(amount float64) {
	if amount <= ba.Balance {
		ba.Balance -= amount
	}
}

func (ba BankAccount) Display() {
	fmt.Println("Account Number:", ba.AccountNumber)
	fmt.Println("Holder Name:", ba.HolderName)
	fmt.Println("Balance:", ba.Balance)
}
