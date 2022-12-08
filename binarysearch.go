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


func main() {
	fmt.Println("### Golang Snippets ### Binary Search")

	// sorted array
	array := []int{2,2,2,2,2,2,2,3,5,5,6,6,7,7,7,7,13,13}

	for index, value := range array {
		fmt.Printf("array[%d] = %d ", index, value)
		fmt.Printf("target: %d, low: %d\n", value, low(array,value))
	}
}
