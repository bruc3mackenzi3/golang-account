package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

const MAX_MESSAGES = 10000

// Receives transactions from input channel, processes them and sends to output
// channel.  Runs until input channel is closes loses output channel when all
// NOTE: Specify channel direction to increase type-safety
func processInput(input <-chan []byte, output chan<- []byte) {
	var depositJSON = make(map[string]string)

	for transaction := range input {
		err := json.Unmarshal(transaction, &depositJSON)
		if err != nil {
			println("Error: input is not valid JSON: ", string(transaction))
			continue // discard transaction if malformed
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

// Reads lines from stdin and sends to input channel.  Closes the channel when
// EOF or an error is encountered.
func readInput(input chan<- []byte) {
	// NOTE: Passing slices through channels is potentially unsafe because
	// slices are pointers to arrays, meaning the passed data can be overwritten.
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			println("Error reading from stdin:", err)
			break
		}
		input <- line
	}
	close(input)
}

// Main runner function.  Runs goroutines to read input and process input, and
// prints results to stdout.
func run() {
	// Large buffered channels to avoid blocking
	var input = make(chan []byte, MAX_MESSAGES)
	var output = make(chan []byte, MAX_MESSAGES)

	go readInput(input)
	go processInput(input, output)
	// range loops over received items in channel until it's closed
	for result := range output {
		fmt.Println(string(result))
	}
}

func main() {
	run()
}
