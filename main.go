/*
Clone repo into go workspace set by GOPATH
go install golang-account
$GOPATH/bin/golang-account
*/

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

func processDeposit(deposit *Deposit) {
	account := GetAccount(deposit.id)
	account.Deposit(deposit)
	fmt.Println(account.balance)
}

func run() {
	decoder := json.NewDecoder(os.Stdin)
	var parsedJSON = make(map[string]string)
	var deposit *Deposit

	for {
		err := decoder.Decode(&parsedJSON)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("Failed to deserialize JSON: ", err)
		}

		deposit = NewDeposit(parsedJSON)
		processDeposit(deposit)
		break
	}
}

func main() {
	run()
}
