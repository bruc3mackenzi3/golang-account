package main

import (
	"reflect"
	"testing"
)

func TestProcessDeposit(t *testing.T) {
	// Test base case
	input := map[string]string{
		"id":          "15887",
		"customer_id": "528",
		"load_amount": "$3318.47",
		"time":        "2000-01-01T00:00:00Z",
	}
	expected := map[string]interface{}{
		"id":          "15887",
		"customer_id": "528",
		"accepted":    true,
	}
	result := processDeposit(input)
	if reflect.DeepEqual(result, expected) != true {
		t.Errorf("Expected response map does not match actual: %v, %v", expected, result)
	}

	// Test duplicate requests are blocked
	expected["accepted"] = false
	result = processDeposit(input)
	if reflect.DeepEqual(result, expected) != true {
		t.Errorf("Expected response map does not match actual: %v, %v", expected, result)
	}
}
