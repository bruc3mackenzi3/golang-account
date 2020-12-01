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
func processInput(depositJSON map[string]string) map[string]interface{} /* string */ {
	// NOTE: For debugging purposes code to return a string with map fields in a
	// particular order is left commented out.  This makes comparison with sample
	// output.txt seamless.

	deposit := NewDeposit(depositJSON)
	account := GetAccount(deposit.customerId, deposit.transTime)
	result := account.DepositFunds(deposit)

	return map[string]interface{}{
		"id":          deposit.id,
		"customer_id": deposit.customerId,
		"accepted":    result,
	}
	// return fmt.Sprintf(`{"id":"%s","customer_id":"%s","accepted":%t}`, deposit.id, deposit.customerId, result)
}

// Main runner function.  Loops over input from stdin, decodes JSON and handles
// errors, and prints results to stdout.
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

		result := processInput(parsedJSON)
		fmt.Println(result)
	}
}

func main() {
	run()
}
