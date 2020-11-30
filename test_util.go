package main

import "time"

const SAMPLE_TIME = "2000-01-01T00:00:00Z"

var sampleDeposit = Deposit{
	id:         "12345",
	customerId: "123",
	loadAmount: 100.0,
	transTime:  ParseTime(SAMPLE_TIME),
}

var sampleLimits = AccountLimits{latestTime: ParseTime(SAMPLE_TIME)}

var sampleAccount = Account{
	balance: 0.0,
	limits:  sampleLimits,
}

// These Get functions return new copies of their respective Struct by making
// use of Go's pass by value property
func GetSampleDeposit() Deposit {
	return sampleDeposit
}

func GetSampleAccountLimits() AccountLimits {
	return sampleLimits
}

func GetSampleAccount() Account {
	return sampleAccount
}

func ParseTime(t string) time.Time {
	timeT, _ := time.Parse(time.RFC3339, t)
	return timeT
}
