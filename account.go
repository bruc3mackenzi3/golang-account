package main

type Account struct {
	balance float64
}

// Deposit funds into Account
func (account *Account) Deposit(amount float64) {
	// TODO: Make amount unsigned
	account.balance = account.balance + amount
}
