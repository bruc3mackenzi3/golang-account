package main

import (
	"testing"
)

func TestGetAccountLimits(t *testing.T) {
	testTime := ParseTime(SAMPLE_TIME)
	expected := AccountLimits{latestTime: testTime}

	result := GetAccountLimits(testTime)
	if expected != *result {
		t.Errorf("AccountLimits returned does not match expected: %v $%v", result, expected)
	}
}

func TestUpdate(t *testing.T) {
	testTime := ParseTime("2020-01-01T00:00:00Z")
	limits := GetSampleAccountLimits()
	limits.latestTime = testTime
	deposit := GetSampleDeposit()
	deposit.transTime = testTime
	deposit.loadAmount = 1000.0

	// 1. Same day
	limits.Update(&deposit)
	if limits.dailyCount != 1 || limits.dailyAmount != 1000.0 || limits.weeklyAmount != 1000.0 {
		t.Errorf("AccountLimits Update failed: %d %f %f", limits.dailyCount, limits.dailyAmount, limits.weeklyAmount)
	}

	limits.Update(&deposit)
	if limits.dailyCount != 2 || limits.dailyAmount != 2000.0 || limits.weeklyAmount != 2000.0 {
		t.Errorf("AccountLimits Update failed: %d %f %f", limits.dailyCount, limits.dailyAmount, limits.weeklyAmount)
	}

	// 2. Different day
	deposit.transTime = ParseTime("2020-01-02T00:00:00Z")
	limits.Update(&deposit)
	if limits.dailyCount != 1 || limits.dailyAmount != 1000.0 || limits.weeklyAmount != 3000.0 {
		t.Errorf("AccountLimits Update failed: %d %f %f", limits.dailyCount, limits.dailyAmount, limits.weeklyAmount)
	}
	deposit.transTime = ParseTime("2020-01-03T00:00:00Z")
	limits.Update(&deposit)
	if limits.dailyCount != 1 || limits.dailyAmount != 1000.0 || limits.weeklyAmount != 4000.0 {
		t.Errorf("AccountLimits Update failed: %d %f %f", limits.dailyCount, limits.dailyAmount, limits.weeklyAmount)
	}

	// 3. Different week
	deposit.transTime = ParseTime("2020-01-23T00:00:00Z")
	limits.Update(&deposit)
	if limits.dailyCount != 1 || limits.dailyAmount != 1000.0 || limits.weeklyAmount != 1000.0 {
		t.Errorf("AccountLimits Update failed: %d %f %f", limits.dailyCount, limits.dailyAmount, limits.weeklyAmount)
	}
	deposit.transTime = ParseTime("2020-02-23T00:00:00Z")
	limits.Update(&deposit)
	if limits.dailyCount != 1 || limits.dailyAmount != 1000.0 || limits.weeklyAmount != 1000.0 {
		t.Errorf("AccountLimits Update failed: %d %f %f", limits.dailyCount, limits.dailyAmount, limits.weeklyAmount)
	}
}

func TestIsDepositLimitReached(t *testing.T) {
	testTime := ParseTime("2000-01-01T00:00:00Z")
	limits := GetSampleAccountLimits() // new AccountLimits with 0s
	limits.latestTime = testTime
	deposit := GetSampleDeposit()
	deposit.loadAmount = 1000.0
	deposit.transTime = testTime

	// 1. Max number of daily transactions
	if limits.IsDepositLimitReached(&deposit) != false {
		t.Error("False positive with account and deposit below limits")
	}

	limits.dailyCount = 2
	if limits.IsDepositLimitReached(&deposit) != false {
		t.Error("False positive with account and deposit below limits")
	}

	limits.dailyCount = 3
	if limits.IsDepositLimitReached(&deposit) != true {
		t.Error("False negative with account and deposit below limits")
	}

	limits.dailyCount = 4
	if limits.IsDepositLimitReached(&deposit) != true {
		t.Error("False negative with account and deposit below limits")
	}

	// 2. Daily transaction amount limit
	limits.dailyCount = 0
	limits.dailyAmount = 0.0
	if limits.IsDepositLimitReached(&deposit) != false {
		t.Error("False positive with account and deposit below limits")
	}
	limits.dailyAmount = DAILY_DEPOSIT_AMOUNT_LIMIT - deposit.loadAmount - 1.0
	if limits.IsDepositLimitReached(&deposit) != false {
		t.Error("False positive with account and deposit below limits")
	}

	limits.dailyAmount = DAILY_DEPOSIT_AMOUNT_LIMIT - deposit.loadAmount
	if limits.IsDepositLimitReached(&deposit) != false {
		t.Error("False positive with account and deposit below limits")
	}

	limits.dailyAmount = DAILY_DEPOSIT_AMOUNT_LIMIT - deposit.loadAmount + 1.0
	if limits.IsDepositLimitReached(&deposit) != true {
		t.Error("False negative with account and deposit below limits")
	}

	limits.dailyAmount = DAILY_DEPOSIT_AMOUNT_LIMIT * 100.0
	if limits.IsDepositLimitReached(&deposit) != true {
		t.Error("False negative with account and deposit below limits")
	}

	// 3. Weekly transaction amount limit
	limits.dailyAmount = 0.0
	limits.weeklyAmount = WEEKLY_DEPOSIT_AMOUNT_LIMIT - deposit.loadAmount - 1.0
	if limits.IsDepositLimitReached(&deposit) != false {
		t.Error("False positive with account and deposit below limits")
	}

	limits.weeklyAmount = WEEKLY_DEPOSIT_AMOUNT_LIMIT - deposit.loadAmount
	if limits.IsDepositLimitReached(&deposit) != false {
		t.Error("False positive with account and deposit below limits")
	}

	limits.weeklyAmount = WEEKLY_DEPOSIT_AMOUNT_LIMIT - deposit.loadAmount + 1.0
	if limits.IsDepositLimitReached(&deposit) != true {
		t.Error("False negative with account and deposit below limits")
	}

	limits.weeklyAmount = WEEKLY_DEPOSIT_AMOUNT_LIMIT * 100.0
	if limits.IsDepositLimitReached(&deposit) != true {
		t.Error("False negative with account and deposit below limits")
	}

	// 4. Everything over
	limits.dailyCount = 4
	limits.dailyAmount = DAILY_DEPOSIT_AMOUNT_LIMIT + 1.0
	limits.weeklyAmount = WEEKLY_DEPOSIT_AMOUNT_LIMIT + 1.0
	if limits.IsDepositLimitReached(&deposit) != true {
		t.Error("False negative with account and deposit below limits")
	}
}

func TestHasSameDay(t *testing.T) {
	testTime := ParseTime("2000-01-01T00:00:00Z")
	limits := GetSampleAccountLimits()
	limits.latestTime = testTime
	deposit := GetSampleDeposit()
	deposit.transTime = testTime

	// Same date and time
	if limits.hasSameDay(&deposit) != true {
		t.Errorf("hasSameDay returns false for same days %v %v", limits.latestTime, deposit.transTime)
	}

	// Same date different time
	deposit.transTime = ParseTime("2000-01-01T12:30:00Z")
	if limits.hasSameDay(&deposit) != true {
		t.Errorf("hasSameDay returns false for same days %v %v", limits.latestTime, deposit.transTime)
	}

	// One day later
	deposit.transTime = ParseTime("2000-01-02T00:00:00Z")
	if limits.hasSameDay(&deposit) != false {
		t.Errorf("hasSameDay returns true for different days %v %v", limits.latestTime, deposit.transTime)
	}

	// One day before
	deposit.transTime = ParseTime("1999-12-31T00:00:00Z")
	if limits.hasSameDay(&deposit) != false {
		t.Errorf("hasSameDay returns true for different days %v %v", limits.latestTime, deposit.transTime)
	}

	// Multiple days later
	deposit.transTime = ParseTime("2000-01-31T00:00:00Z")
	if limits.hasSameDay(&deposit) != false {
		t.Errorf("hasSameDay returns true for different days %v %v", limits.latestTime, deposit.transTime)
	}

	// One month later
	deposit.transTime = ParseTime("2000-02-01T00:00:00Z")
	if limits.hasSameDay(&deposit) != false {
		t.Errorf("hasSameDay returns true for different days %v %v", limits.latestTime, deposit.transTime)
	}

	// Multiple months later
	deposit.transTime = ParseTime("2000-12-01T00:00:00Z")
	if limits.hasSameDay(&deposit) != false {
		t.Errorf("hasSameDay returns true for different days %v %v", limits.latestTime, deposit.transTime)
	}

	// One year later
	deposit.transTime = ParseTime("2001-01-01T00:00:00Z")
	if limits.hasSameDay(&deposit) != false {
		t.Errorf("hasSameDay returns true for different days %v %v", limits.latestTime, deposit.transTime)
	}

	// Multiple years later
	deposit.transTime = ParseTime("2009-01-01T00:00:00Z")
	if limits.hasSameDay(&deposit) != false {
		t.Errorf("hasSameDay returns true for different days %v %v", limits.latestTime, deposit.transTime)
	}
}

func TestHasSameWeek(t *testing.T) {
	testTime := ParseTime("2020-01-01T00:00:00Z")
	limits := GetSampleAccountLimits()
	limits.latestTime = testTime
	deposit := GetSampleDeposit()
	deposit.transTime = testTime

	// Same date and time
	if limits.hasSameWeek(&deposit) != true {
		t.Errorf("hasSameWeek returns false for same weeks %v %v", limits.latestTime, deposit.transTime)
	}

	// Different day same week
	deposit.transTime = ParseTime("2020-01-02T00:00:00Z")
	if limits.hasSameWeek(&deposit) != true {
		t.Errorf("hasSameWeek returns false for same weeks %v %v", limits.latestTime, deposit.transTime)
	}

	// Different year same week
	limits.latestTime = ParseTime("2019-12-31T00:00:00Z")
	if limits.hasSameWeek(&deposit) != true {
		t.Errorf("hasSameWeek returns false for same weeks %v %v", limits.latestTime, deposit.transTime)
	}

	// Different year different week
	limits.latestTime = ParseTime("2019-12-28T00:00:00Z")
	if limits.hasSameWeek(&deposit) != false {
		t.Errorf("hasSameWeek returns true for different weeks %v %v", limits.latestTime, deposit.transTime)
	}

	// Different year different week by 1 second
	limits.latestTime = ParseTime("2019-12-28T23:59:59Z")
	if limits.hasSameWeek(&deposit) != false {
		t.Errorf("hasSameWeek returns true for different weeks %v %v", limits.latestTime, deposit.transTime)
	}

	// Much different year
	limits.latestTime = ParseTime("2000-01-01T00:00:00Z")
	if limits.hasSameWeek(&deposit) != false {
		t.Errorf("hasSameWeek returns true for different weeks %v %v", limits.latestTime, deposit.transTime)
	}
}
