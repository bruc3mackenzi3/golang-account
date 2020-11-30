package main

import "testing"

func TestNewDeposit(t *testing.T) {
	input := map[string]string{
		"id":          "11111",
		"customer_id": "222",
		"load_amount": "$3318.47",
		"time":        "2000-01-01T00:00:00Z",
	}
	expected := Deposit{
		id:          "11111",
		customer_id: "222",
		load_amount: 3318.47,
		time:        "2000-01-01T00:00:00Z",
	}
	result := NewDeposit(input)
	if *result != expected {
		t.Errorf("Expected Deposit struct does not match actual: %v, %v", expected, result)
	}
}

type testLoadAmountData struct {
	input    string
	expected float64
}

func TestParseLoadAmount(t *testing.T) {
	validTests := []testLoadAmountData{
		testLoadAmountData{input: "$1000.00", expected: 1000.00},
		testLoadAmountData{input: "$999.99", expected: 999.99},
		testLoadAmountData{input: "$3318.47", expected: 3318.47},
		testLoadAmountData{input: "$1.01", expected: 1.01},
		testLoadAmountData{input: "$1.00", expected: 1.00},
		testLoadAmountData{input: "$0.01", expected: 0.01},
	}

	for _, test := range validTests {
		result, err := parseLoadAmount(test.input)
		if result != test.expected {
			t.Errorf("Got %f, expected %f", 0.0, 0.0)
		}
		if err != nil {
			t.Errorf("Got error on valid input %s: %v", test.input, err)
		}
	}

	// TODO: Add testing for NaN and other edge cases
	errorTests := []string{
		"100.00",
		"0.99",
		"$0.00",
		"$-0.01",
		"$-1.00",
	}
	for _, input := range errorTests {
		result, err := parseLoadAmount(input)
		if result != 0.0 {
			t.Errorf("Got non-zero result on error.  Got %f, expected 0.0", result)
		}
		if err == nil {
			t.Errorf("No error on bad input %s", input)
		}
	}
}

func TestIsDepositProcessed(t *testing.T) {
	dep := Deposit{
		id:          "13579",
		customer_id: "135",
		load_amount: 1000.00,
		time:        "2000-01-01T10:00:00Z",
	}
	result := dep.IsDepositProcessed()
	if result != false {
		t.Error("IsDepositProcessed returned true, expected false")
	}

	dep.RecordDeposit()
	result = dep.IsDepositProcessed()
	if result != true {
		t.Error("IsDepositProcessed returned false, expected true")
	}

	// TODO: Add more tests to ensure robustness
}
