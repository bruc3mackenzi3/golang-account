package main

import "time"

type Account struct {
	balance float64
	limits  AccountLimits
}

var accounts = make(map[string]Account)

// Retrieve the Account associated with id, or return a newly created one if it
// doesn't exist.
// Note that a newly created Account will have its time set from the current
// transaction, while an existing Account's time was previously set by the
// previous transaction time.
func GetAccount(id string, transTime time.Time) *Account {
	_, found := accounts[id]
	if found == false {
		accounts[id] = Account{
			balance: 0.0,
			limits:  *GetAccountLimits(transTime),
		}
	}
	account, _ := accounts[id]
	return &account
}

// Deposit funds into Account
// Returns true if Deposit is accepted, false otherwise
// NOTE: This function is NOT thread safe
func (account *Account) DepositFunds(deposit *Deposit) bool {
	if deposit.IsDepositProcessed() == true {
		return false
	}

	// Reject deposit if limit on Account has been reached
	if account.limits.IsDepositLimitReached(deposit) == true {
		return false
	}

	// Add funds to account and commit transaction!
	account.balance = account.balance + deposit.loadAmount
	account.limits.Update(deposit)
	deposit.Record()
	return true
}
