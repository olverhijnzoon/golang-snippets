package main

import (
	"fmt"
	"time"
)

const (
	year         = 2022
	firstAdvent  = 27
	secondAdvent = 7
	thirdAdvent  = 14
	fourthAdvent = 21
	dateFormat   = "2006-01-02"
	hours        = 0
	minutes      = 0
	seconds      = 0
	nanoSeconds  = 0
	addYear      = 0
	addMonth     = 0
	addDay       = 1
)

var location = time.UTC

func main() {

	fmt.Println("# Golang Snippets")
	fmt.Println("## Advent")

	// Calculate the date for the first Sunday of Advent
	firstSunday := time.Date(year, time.November, firstAdvent, hours, minutes, seconds, nanoSeconds, location)
	for firstSunday.Weekday() != time.Sunday {
		firstSunday = firstSunday.AddDate(addYear, addMonth, addDay)
	}

	// Print the dates for the four Advent Sundays
	fmt.Println("First Advent Sunday:", firstSunday.Format(dateFormat))
	fmt.Println("Second Advent Sunday:", firstSunday.AddDate(addYear, addMonth, secondAdvent).Format(dateFormat))
	fmt.Println("Third Advent Sunday:", firstSunday.AddDate(addYear, addMonth, thirdAdvent).Format(dateFormat))
	fmt.Println("Fourth Advent Sunday:", firstSunday.AddDate(addYear, addMonth, fourthAdvent).Format(dateFormat))
}
