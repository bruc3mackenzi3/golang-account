package main

type Account struct {
	balance float64
}

var accounts = make(map[string]Account)

// Retrieve the Account associated with id, or return a newly created one if it
// doesn't exist
func GetAccount(id string) *Account {
	_, found := accounts[id]
	if found == false {
		accounts[id] = Account{0.0}
	}
	account, _ := accounts[id]
	return &account
}

// Deposit funds into Account
func (account *Account) Deposit(deposit *Deposit) {
	account.balance = account.balance + deposit.load_amount
}
