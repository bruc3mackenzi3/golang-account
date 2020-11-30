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

// Perform critical business logic checking if velocity limit on Account will
// be reached with Deposit.  Return true if one of the limits has been reached,
// false otherwise.
func (limit *AccountLimits) IsDepositLimitReached(deposit *Deposit) bool {
	if limit.hasSameDay(deposit) {
		if limit.dailyCount >= DAILY_DEPOSIT_COUNT_LIMIT {
			return true
		}
	}
	return false
}

// Update account limit counts when a deposit is processed.  Assumes
// IsDepositLimitReached() was called and returned false.
func (limit *AccountLimits) Update(deposit *Deposit) {

	if limit.hasSameDay(deposit) {
		// If day is the same update daily transaction count and amount
		// deposited
		limit.dailyCount++
		limit.dailyAmount += deposit.loadAmount
	} else {
		// If day is different update the day and reset counts
		limit.dailyCount = 0
		limit.dailyAmount = 0.0
		limit.latestTime = deposit.transTime
	}
}

// Helper function to determine if previous and current deposits fall on the
// same day
func (limit *AccountLimits) hasSameDay(deposit *Deposit) bool {
	prev := limit.latestTime
	curr := deposit.transTime
	if prev.Year() == curr.Year() && prev.YearDay() == curr.YearDay() {
		return true
	} else {
		return false
	}
}
