package quicksort

import (
	"math"
	"reflect"
	"testing"
)

func TestQuicksort(t *testing.T) {
	testCases := []struct {
		input    []int
		expected []int
	}{
		{[]int{5, 2, 6, 1, 3, -5}, []int{-5, 1, 2, 3, 5, 6}},
		{[]int{5, 2, 6, 1, 3}, []int{1, 2, 3, 5, 6}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{3, 3, 3}, []int{3, 3, 3}},
		{[]int{}, []int{}},
		{[]int{math.MaxInt32, math.MaxInt32, math.MaxInt32}, []int{math.MaxInt32, math.MaxInt32, math.MaxInt32}},
		{[]int{math.MinInt32, math.MinInt32, math.MinInt32}, []int{math.MinInt32, math.MinInt32, math.MinInt32}},
		{[]int{math.MaxInt32, math.MinInt32, 100, -100}, []int{math.MinInt32, -100, 100, math.MaxInt32}},
		{[]int{10, 20, 30, 40, 50, 60}, []int{10, 20, 30, 40, 50, 60}},
		{[]int{-10, -20, -30, -40, -50, -60}, []int{-60, -50, -40, -30, -20, -10}},
		{[]int{0, 10, -10, 20, -20, 30, -30}, []int{-30, -20, -10, 0, 10, 20, 30}},
	}

	for _, tc := range testCases {
		result := Quicksortalgo(tc.input)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("quicksort(%v) = %v, expected %v", tc.input, result, tc.expected)
		}
	}
}
