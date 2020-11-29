package main

import (
	"errors"
	"log"
	"os"
	"strconv"
)

type Deposit struct {
	id          string
	customer_id string
	load_amount float64
	time        string
}

func NewDeposit(input map[string]string) *Deposit {
	dep := Deposit{}
	var err error

	dep.id = input["id"]
	dep.customer_id = input["customer_id"]
	dep.load_amount, err = parseLoadAmount(input["load_amount"])
	dep.time = input["time"]

	if err != nil {
		log.Fatal("Failed to parse load_amount", input["load_amount"], "from input", input, "error:", err)
		os.Exit(1)
	}
	return &dep
}

// Parse the raw load amount string into a floating point number.  Returns an
// error if parsing failed or the value is less than or equal to 0.0
func parseLoadAmount(input string) (float64, error) {
	if input[0] != '$' {
		return 0.0, errors.New("load_amount missing leading '$'")
	}
	amount, err := strconv.ParseFloat(input[1:], 64)
	if err != nil {
		return 0.0, err
	}
	if amount <= 0.0 {
		return 0.0, errors.New("load_amount must be greater then 0.0")
	}
	return amount, nil
}
