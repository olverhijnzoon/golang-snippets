package quicksort

import (
	"reflect"
	"testing"
)

func TestQuicksort(t *testing.T) {
	testCases := []struct {
		input    []int
		expected []int
	}{
		{[]int{5, 2, 6, 1, 3}, []int{1, 2, 3, 5, 6}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{3, 3, 3}, []int{3, 3, 3}},
		{[]int{}, []int{}},
	}

	for _, tc := range testCases {
		result := Quicksortalgo(tc.input)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("quicksort(%v) = %v, expected %v", tc.input, result, tc.expected)
		}
	}
}
