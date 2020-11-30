package main

import (
	"testing"
)

type testGetAmountData struct {
	input    string
	expected Account
}

func TestGetAccount(t *testing.T) {
	rawTime := "2000-01-01T00:00:00Z"

	expectedAccount := Account{
		balance: 0.0,
		limits:  *GetAccountLimits(ParseTime(rawTime)),
	}
	tests := []testGetAmountData{
		testGetAmountData{input: "123", expected: expectedAccount},
		testGetAmountData{input: "123", expected: expectedAccount},
		testGetAmountData{input: "456", expected: expectedAccount},
	}

	for _, testData := range tests {
		result := GetAccount(testData.input, ParseTime(rawTime))
		if testData.expected != *result {
			t.Errorf("Expected Account struct does not match actual: %v, %v", testData.expected, result)
		}
	}

	// TODO: Test after balance is changed
}

func TestDeposit(*testing.T) {

}
