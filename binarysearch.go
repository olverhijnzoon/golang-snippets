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

// with binary search give back the number of occurrences of the target value in the array
func count(arr []int, target int) int {
	lowIndex := low(arr, target)
	highIndex := high(arr, target)
	return highIndex - lowIndex + 1
}

// with binary search give back the first and last indexes and the count of the target value in the array
func firstLastCount(arr []int, target int) (int, int, int) {
	lowIndex := low(arr, target)
	highIndex := high(arr, target)
	return lowIndex, highIndex, highIndex - lowIndex + 1
}

func main() {
	fmt.Println("### Golang Snippets ### Binary Search")

	// sorted array
	array := []int{2,2,2,2,2,2,2,3,5,5,6,6,7,7,7,7,13,13}

	// use of map to keep track of the mentioned target values
	targets := map[int]bool{}

	for index, value := range array {	
		if _, ok := targets[value]; !ok {
			targets[value] = true
			lowIndex, highIndex, count := firstLastCount(array, value)
			fmt.Printf("Result: target %d, low %d, high %d, count %d\n", value, lowIndex, highIndex, count)
			fmt.Printf("Check: array[%d] = %d \n", index, value)
		}
	}
}
