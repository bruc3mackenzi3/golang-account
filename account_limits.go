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
	// NOTE: Check daily limit against both load amount alone and in
	// combination with daily amount
	if deposit.loadAmount > DAILY_DEPOSIT_AMOUNT_LIMIT {
		return true
	}
	if limit.hasSameDay(deposit) {
		if limit.dailyCount >= DAILY_DEPOSIT_COUNT_LIMIT {
			return true
		}
		if limit.dailyAmount+deposit.loadAmount > DAILY_DEPOSIT_AMOUNT_LIMIT {
			return true
		}
	}
	if limit.hasSameWeek(deposit) {
		if limit.weeklyAmount+deposit.loadAmount > WEEKLY_DEPOSIT_AMOUNT_LIMIT {
			return true
		}
	}
	return false
}

// Update account limit counts when a deposit is processed.  Assumes
// IsDepositLimitReached() was called and returned false.
func (limit *AccountLimits) Update(deposit *Deposit) {
	if limit.hasSameDay(deposit) {
		// If day is the same update daily transaction count and daily deposit
		// amount
		limit.dailyCount++
		limit.dailyAmount = limit.dailyAmount + deposit.loadAmount
	} else {
		// If day is different reset counts
		limit.dailyCount = 1
		limit.dailyAmount = deposit.loadAmount
	}

	// NOTE: dailyCount was updated above
	if limit.hasSameWeek(deposit) {
		// If week is the same update weekly deposit amount
		limit.weeklyAmount = limit.weeklyAmount + deposit.loadAmount
	} else {
		limit.weeklyAmount = deposit.loadAmount
	}

	// Finally update the latest transaction time
	limit.latestTime = deposit.transTime
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

// Helper function to determine if previous and current deposits fall in the
// same week
// NOTE: assumes transactions are always processed in ascending order
func (limit *AccountLimits) hasSameWeek(deposit *Deposit) bool {
	prev := limit.latestTime
	curr := deposit.transTime

	prevYear, prevWeek := prev.ISOWeek()
	currYear, currWeek := curr.ISOWeek()
	if prevYear == currYear && prevWeek == currWeek {
		return true
	} else {
		return false
	}
}
