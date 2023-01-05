package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("# Golang Snippets")
	fmt.Println("## Maps")

	/*
		This code creates maps with 1000, 100,000, and 10,000,000 elements, and then measures the time it takes to look up an element in each of these maps 4206912 times. It then calculates the average time for each map size by dividing the total time by the number of measurements. Because the time complexity of map lookups in Go is constant, you should see that the average time it takes to perform the lookups is roughly the same regardless of the size of the map.

		Map lookups in Go have constant-time complexity because they are implemented using hash tables. A hash table is a data structure that stores key-value pairs and allows for efficient insertion, deletion, and lookup of values based on their keys.

		In a hash table, the keys are hashed to produce a numeric value, which is then used to determine the location in the table where the value should be stored. When looking up a value in the table, the key is hashed again and the resulting numeric value is used to find the location in the table where the value is stored. Because the hash function is designed to evenly distribute the keys across the table, the time it takes to look up a value in a hash table is approximately constant, regardless of the size of the table.

		This is why map lookups in Go have constant-time complexity: because they are implemented using hash tables, the time it takes to look up a value in a map is approximately constant, regardless of the size of the map.
	*/

	// Measure the time it takes to look up an element in a map with 1000 elements, 4206912 times.
	m := make(map[int]int)
	for i := 0; i < 1000; i++ {
		m[i] = i * i
	}
	totalTime := 0 * time.Millisecond
	for i := 0; i < 4206912; i++ {
		start := time.Now()
		_ = m[999]
		elapsed := time.Since(start)
		totalTime += elapsed
	}
	avgTime := totalTime / 4206912

	fmt.Printf("\nAverage time to look up element in map with 1000 elements: %v\n", avgTime)

	// Measure the time it takes to look up an element in a map with 100,000 elements, 4206912 times.
	m = make(map[int]int)
	for i := 0; i < 100000; i++ {
		m[i] = i * i
	}
	totalTime = 0 * time.Millisecond
	for i := 0; i < 4206912; i++ {
		start := time.Now()
		_ = m[99999]
		elapsed := time.Since(start)
		totalTime += elapsed
	}
	avgTime = totalTime / 4206912
	fmt.Printf("\nAverage time to look up element in map with 100,000 elements: %v\n", avgTime)

	// Measure the time it takes to look up an element in a map with 10,000,000 elements, 4206912 times.
	m = make(map[int]int)
	for i := 0; i < 10000000; i++ {
		m[i] = i * i
	}
	totalTime = 0 * time.Millisecond
	for i := 0; i < 4206912; i++ {
		start := time.Now()
		_ = m[9999999]
		elapsed := time.Since(start)
		totalTime += elapsed
	}
	avgTime = totalTime / 4206912
	fmt.Printf("\nAverage time to look up element in map with 10,000,000 elements: %v\n", avgTime)
}
