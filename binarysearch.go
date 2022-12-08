package main

import "fmt"

// with binary search give back index of the first element in the sorted array that is greater or equal to target
func low(arr []int, target int) int{
	low := 0
	high:= len(arr)-1

	for low <= high {
		mid := (low+high)/2
		if arr[mid]>=target{
			high = mid -1
		} else {
			low = mid +1
		}
	}

	return low
}

// with binary search give back index of the last element in the sorted array that is less or equal to target
func high(arr []int, target int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2
		if arr[mid] <= target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return high
}

func main() {
	fmt.Println("### Golang Snippets ### Binary Search")

	// sorted array
	array := []int{2,2,2,2,2,2,2,3,5,5,6,6,7,7,7,7,13,13}

	for index, value := range array {
		fmt.Printf("array[%d] = %d ", index, value)
		fmt.Printf("low: %d, high: %d\n", low(array, value), high(array, value))
	}
}
