package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

const MAX_WORKERS = 100

// Given parsed JSON input process transaction and return a map with accepted
// status.
func processInput(input chan []byte, output chan []byte) {
	var depositJSON = make(map[string]string)
	for transaction := range input {
		err := json.Unmarshal(transaction, &depositJSON)
		if err != nil {
			println("Error: input is not valid JSON: ", string(transaction))
			continue
		}

		deposit := NewDeposit(depositJSON)
		account := GetAccount(deposit.customerId, deposit.transTime)
		result := account.DepositFunds(deposit)

		result_json, err := json.Marshal(map[string]interface{}{
			"id":          deposit.id,
			"customer_id": deposit.customerId,
			"accepted":    result,
		})
		output <- result_json
	}
	close(output)
}

func read_input(input chan []byte, output chan []byte) {
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			println("Failed to deserialize JSON: ", err)
			continue
		}
		input <- line
	}
	close(input)
}

// Main runner function.  Loops over input from stdin, decodes JSON and handles
// errors, and prints results to stdout.
func run() {
	var input = make(chan []byte, MAX_WORKERS)
	var output = make(chan []byte, MAX_WORKERS)

	go read_input(input, output)
	go processInput(input, output)
	for result := range output {
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
