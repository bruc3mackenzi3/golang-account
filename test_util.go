package main

import "time"

func ParseTime(t string) time.Time {
	timeT, _ := time.Parse(time.RFC3339, t)
	return timeT
}
