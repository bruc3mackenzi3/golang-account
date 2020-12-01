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

// Given parsed JSON input process transaction and return a map with accepted
// status.
func processDeposit(depositJSON map[string]string) map[string]interface{} {
	deposit := NewDeposit(depositJSON)
	account := GetAccount(deposit.id, deposit.transTime)
	result := account.DepositFunds(deposit)
	return map[string]interface{}{
		"id":          deposit.id,
		"customer_id": deposit.customerId,
		"accepted":    result,
	}
}

// Main runner function.  Loops over input from stdin and prints corresponding
// results to stdout.
func run() {
	decoder := json.NewDecoder(os.Stdin)
	var parsedJSON = make(map[string]string)

	for {
		err := decoder.Decode(&parsedJSON)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("Failed to deserialize JSON: ", err)
		}

		result := processDeposit(parsedJSON)
		output, _ := json.Marshal(result)
		fmt.Println(string(output))
	}
}

func main() {
	run()
}
