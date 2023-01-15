package main

import (
	"fmt"
	"time"
)

const (
	year                                                 = 2022
	firstAdvent, secondAdvent, thirdAdvent, fourthAdvent = 27, 7, 14, 21
	dateFormat                                           = "2006-01-02"
	hours, minutes, seconds, nanoSeconds                 = 0, 0, 0, 0
	addYear, addMonth, addDay                            = 0, 0, 1
)

var location = time.UTC

func calculateAdventSundays(year int) (first, second, third, fourth time.Time) {
	// Calculate the date for the first Sunday of Advent
	first = time.Date(year, time.November, firstAdvent, hours, minutes, seconds, nanoSeconds, location)
	for first.Weekday() != time.Sunday {
		first = first.AddDate(addYear, addMonth, addDay)
	}

	// Calculate the dates for the remaining Advent Sundays
	second = first.AddDate(addYear, addMonth, secondAdvent)
	third = first.AddDate(addYear, addMonth, thirdAdvent)
	fourth = first.AddDate(addYear, addMonth, fourthAdvent)
	return
}

func main() {
	fmt.Println("# Golang Snippets")
	fmt.Println("## Advent")

	first, second, third, fourth := calculateAdventSundays(year)

	// Print the dates for the four Advent Sundays
	fmt.Println("First Advent Sunday:", first.Format(dateFormat))
	fmt.Println("Second Advent Sunday:", second.Format(dateFormat))
	fmt.Println("Third Advent Sunday:", third.Format(dateFormat))
	fmt.Println("Fourth Advent Sunday:", fourth.Format(dateFormat))
}
