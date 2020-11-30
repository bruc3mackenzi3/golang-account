package main

import (
	"testing"
)

func TestIsDepositLimitReached(t *testing.T) {
	testTime := ParseTime("2000-01-01T00:00:00Z")
	limits := GetSampleAccountLimits()
	limits.latestTime = testTime
	deposit := GetSampleDeposit()
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
	limits.dailyAmount = DAILY_DEPOSIT_AMOUNT_LIMIT - 1.0
	if limits.IsDepositLimitReached(&deposit) != false {
		t.Error("False positive with account and deposit below limits")
	}

	// limits.dailyAmount = DAILY_DEPOSIT_AMOUNT_LIMIT
	// if limits.IsDepositLimitReached(&deposit) != true {
	// 	t.Error("False negative with account and deposit below limits")
	// }

	// limits.dailyAmount = DAILY_DEPOSIT_AMOUNT_LIMIT + 1.0
	// if limits.IsDepositLimitReached(&deposit) != true {
	// 	t.Error("False negative with account and deposit below limits")
	// }

	// limits.dailyAmount = DAILY_DEPOSIT_AMOUNT_LIMIT * 100.0
	// if limits.IsDepositLimitReached(&deposit) != true {
	// 	t.Error("False negative with account and deposit below limits")
	// }

	// 3. Weekly transaction amount limit
	limits.dailyAmount = 0.0
	limits.weeklyAmount = WEEKLY_DEPOSIT_AMOUNT_LIMIT - 1.0
	if limits.IsDepositLimitReached(&deposit) != false {
		t.Error("False positive with account and deposit below limits")
	}

	// limits.weeklyAmount = WEEKLY_DEPOSIT_AMOUNT_LIMIT
	// if limits.IsDepositLimitReached(&deposit) != true {
	// 	t.Error("False negative with account and deposit below limits")
	// }

	// limits.weeklyAmount = WEEKLY_DEPOSIT_AMOUNT_LIMIT + 1.0
	// if limits.IsDepositLimitReached(&deposit) != true {
	// 	t.Error("False negative with account and deposit below limits")
	// }

	// limits.weeklyAmount = WEEKLY_DEPOSIT_AMOUNT_LIMIT * 100.0
	// if limits.IsDepositLimitReached(&deposit) != true {
	// 	t.Error("False negative with account and deposit below limits")
	// }
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

	// MUltiple months later
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
