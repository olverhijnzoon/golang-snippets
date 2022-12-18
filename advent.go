package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("# Golang Snippets")
	fmt.Println("## Advent")

	year := 2022 // specify the year for which to calculate Advent Sundays

	// Calculate the date for the first Sunday of Advent
	firstSunday := time.Date(year, time.November, 27, 0, 0, 0, 0, time.UTC)
	for firstSunday.Weekday() != time.Sunday {
		firstSunday = firstSunday.AddDate(0, 0, 1)
	}

	// Print the dates for the four Advent Sundays
	fmt.Println("First Advent Sunday:", firstSunday.Format("2006-01-02"))
	fmt.Println("Second Advent Sunday:", firstSunday.AddDate(0, 0, 7).Format("2006-01-02"))
	fmt.Println("Third Advent Sunday:", firstSunday.AddDate(0, 0, 14).Format("2006-01-02"))
	fmt.Println("Fourth Advent Sunday:", firstSunday.AddDate(0, 0, 21).Format("2006-01-02"))
}
