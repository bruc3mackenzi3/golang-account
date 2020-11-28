/*
Clone repo into go workspace set by GOPATH
go install golang-account
$GOPATH/bin/golang-account
*/

package main

import "fmt"

func main() {
	account := &Account{0.0}
	account.Deposit(100.50)
	fmt.Printf("Account balance is %f\n", account.balance)
}
