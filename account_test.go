package main

import "testing"

type testGetAmountData struct {
	input    string
	expected Account
}

func TestGetAccount(t *testing.T) {
	tests := []testGetAmountData{
		testGetAmountData{input: "123", expected: Account{balance: 0.0}},
		testGetAmountData{input: "123", expected: Account{balance: 0.0}},
		testGetAmountData{input: "456", expected: Account{balance: 0.0}},
	}

	for _, test := range tests {
		result := GetAccount(test.input)
		if test.expected != *result {
			t.Errorf("Expected Account struct does not match actual: %v, %v", test.expected, result)
		}
	}
}
