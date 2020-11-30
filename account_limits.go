package main

import (
	"time"
)

// time formats: https://golang.org/pkg/time/#pkg-constants

const DAILY_DEPOSIT_COUNT_LIMIT = 3
const DAILY_DEPOSIT_AMOUNT_LIMIT = 5000.0
const WEEKLY_DEPOSIT_AMOUNT_LIMIT = 20000.0

type AccountLimits struct {
	latestTime   time.Time
	dailyCount   int
	dailyAmount  float64
	weeklyAmount float64
}

func GetAccountLimits(transTime time.Time) *AccountLimits {
	return &AccountLimits{
		latestTime:   transTime,
		dailyCount:   0,
		dailyAmount:  0.0,
		weeklyAmount: 0.0,
	}
}
