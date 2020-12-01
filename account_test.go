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
		if testData.expected.balance != result.balance {
			t.Errorf("Expected Account struct does not match actual: %+v, %+v", &testData.expected.balance, result.balance)
		} else if testData.expected.limits != result.limits {
			t.Errorf("Expected Account limits does not match actual")
		}
	}

	// TODO: Test after balance is changed
}

func TestDepositFunds(t *testing.T) {
	id := "709"
	testTime := ParseTime("2020-01-01T00:00:00Z")

	acc := GetAccount(id, testTime)

	dep := GetSampleDeposit()
	dep.transTime = testTime
	dep.loadAmount = 3000.0

	// First deposit
	if acc.DepositFunds(&dep) != true {
		t.Errorf("Deposit failed")
	}

	// Fail on duplicate id
	if acc.DepositFunds(&dep) != false {
		t.Errorf("Deposit did NOT fail")
	}

	// Fail on daily amount limit
	dep.id = "98765"
	if acc.DepositFunds(&dep) != false {
		t.Errorf("Deposit did NOT fail")
	}

	// Fail on single large deposit
	// dep1 := Deposit{id: "298765", customerId: "385", loadAmount: 2200.0, transTime: testTime}
	// acc.DepositFunds(&dep1)

	dep2 := Deposit{id: "75832", customerId: "385", loadAmount: 6000.0, transTime: ParseTime("2020-01-02T00:00:00Z")}
	if acc.DepositFunds(&dep2) != false {
		t.Errorf("Deposit did NOT fail")
	}
}
