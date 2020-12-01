package main

import (
	"errors"
	"log"
	"os"
	"strconv"
	"time"
)

// Deposit represents a deposit transaction attempting to be processed
type Deposit struct {
	id         string
	customerId string
	loadAmount float64
	transTime  time.Time
}

// Go implementation of a set data structure.  Stores a history of deposits to
// ensure duplicates aren't processed.
// Interface with depositHistory through IsDepositProcessed() and Record().
var depositHistory = make(map[string]bool)

// Creates a new Deposit struct from a Map.  Performs validation on the input
// fields, logging and exiting on error.
func NewDeposit(input map[string]string) *Deposit {
	// Datetime strings are expected in the format represented by time.RFC3339
	// time formats: https://golang.org/pkg/time/#pkg-constants
	transTime, err := time.Parse(time.RFC3339, input["time"])
	if err != nil {
		log.Fatal("Failed to parse time", input["time"], "from input.  Time must be in RFC3339 format. Error:", err)
		os.Exit(1)
	}

	loadAmount, err := parseLoadAmount(input["load_amount"])
	if err != nil {
		log.Fatal("Failed to parse load_amount", input["load_amount"], "from input", input, "error:", err)
		os.Exit(1)
	}

	dep := Deposit{
		id:         input["id"],
		customerId: input["customer_id"],
		loadAmount: loadAmount,
		transTime:  transTime,
	}
	return &dep
}

// Parse the raw load amount string into a floating point number.  Returns an
// error if parsing failed or the value is less than or equal to 0.0.
func parseLoadAmount(input string) (float64, error) {
	if input[0] != '$' {
		return 0.0, errors.New("load_amount missing leading '$'")
	}
	amount, err := strconv.ParseFloat(input[1:], 64)
	if err != nil {
		return 0.0, err
	}
	if amount <= 0.0 {
		return 0.0, errors.New("load_amount must be a positive number")
	}
	return amount, nil
}

// Returns true if a deposit with the same ID has already been processed, false
// otherwise.
func (deposit *Deposit) IsDepositProcessed() bool {
	_, exists := depositHistory[deposit.id+deposit.customerId]
	return exists
}

// Stores a record of Deposit indicating it's been processed
// Precondition: All checks have passed and funds have been deposited into
// account.
func (deposit *Deposit) Record() {
	depositHistory[deposit.id+deposit.customerId] = true
}
