package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

// Given parsed JSON input process transaction and return a map with accepted
// status.
func processInput(input []byte) []byte {
	var depositJSON = make(map[string]string)
	err := json.Unmarshal(input, &depositJSON)
	if err != nil {
		return []byte("Error: input is not valid JSON")
	}

	deposit := NewDeposit(depositJSON)
	account := GetAccount(deposit.customerId, deposit.transTime)
	result := account.DepositFunds(deposit)

	output, err := json.Marshal(map[string]interface{}{
		"id":          deposit.id,
		"customer_id": deposit.customerId,
		"accepted":    result,
	})
	return output
}

// Main runner function.  Loops over input from stdin, decodes JSON and handles
// errors, and prints results to stdout.
func run() {
	reader := bufio.NewReader(os.Stdin)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("Failed to deserialize JSON: ", err)
		}
		result := processInput(line)
		fmt.Println(string(result))
	}

	// decoder := json.NewDecoder(os.Stdin)
	// var parsedJSON = make(map[string]string)

	// for {
	// 	err := decoder.Decode(&parsedJSON)
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err != nil {
	// 		log.Fatal("Failed to deserialize JSON: ", err)
	// 	}

	// 	result := processInput(parsedJSON)
	// 	output, _ := json.Marshal(result)
	// 	fmt.Println(string(output))
	// }
}

func main() {
	run()
}
