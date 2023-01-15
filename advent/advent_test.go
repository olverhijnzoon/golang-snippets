package main

import (
	"testing"
	"time"
)

var nobelYears = []int{1905, 1915, 1921}
var expectedAdventSundays = map[int]map[int]time.Time{
	1905: {
		1: time.Date(1905, time.December, 3, 0, 0, 0, 0, time.UTC),
		2: time.Date(1905, time.December, 10, 0, 0, 0, 0, time.UTC),
		3: time.Date(1905, time.December, 17, 0, 0, 0, 0, time.UTC),
		4: time.Date(1905, time.December, 24, 0, 0, 0, 0, time.UTC),
	},
	1915: {
		1: time.Date(1915, time.November, 28, 0, 0, 0, 0, time.UTC),
		2: time.Date(1915, time.December, 5, 0, 0, 0, 0, time.UTC),
		3: time.Date(1915, time.December, 12, 0, 0, 0, 0, time.UTC),
		4: time.Date(1915, time.December, 19, 0, 0, 0, 0, time.UTC),
	},
	1921: {
		1: time.Date(1921, time.November, 27, 0, 0, 0, 0, time.UTC),
		2: time.Date(1921, time.December, 4, 0, 0, 0, 0, time.UTC),
		3: time.Date(1921, time.December, 11, 0, 0, 0, 0, time.UTC),
		4: time.Date(1921, time.December, 18, 0, 0, 0, 0, time.UTC),
	},
}

func TestCalculateAdventSundays(t *testing.T) {
	for _, year := range nobelYears {
		first, second, third, fourth := calculateAdventSundays(year)
		if !first.Equal(expectedAdventSundays[year][1]) {
			t.Errorf("Incorrect date for first Advent Sunday of %d, got: %v, want: %v", year, first, expectedAdventSundays[year][1])
		}
		if !second.Equal(expectedAdventSundays[year][2]) {
			t.Errorf("Incorrect date for second Advent Sunday of %d, got: %v, want: %v", year, second, expectedAdventSundays[year][2])
		}
		if !third.Equal(expectedAdventSundays[year][3]) {
			t.Errorf("Incorrect date for third Advent Sunday of %d, got: %v, want: %v", year, third, expectedAdventSundays[year][3])
		}
		if !fourth.Equal(expectedAdventSundays[year][4]) {
			t.Errorf("Incorrect date for fourth Advent Sunday of %d, got: %v, want: %v", year, fourth, expectedAdventSundays[year][4])
		}
	}
}
