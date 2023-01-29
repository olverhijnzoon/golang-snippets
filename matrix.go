package main

import (
	"fmt"
)

func main() {

	fmt.Println("# Golang Snippets")
	fmt.Println("## Matrix")

	// Initialize two matrices
	matrixA := [][]int{{1, 2}, {3, 4}}
	matrixB := [][]int{{5, 6}, {7, 8}}

	// Get the number of rows and columns for both matrices
	rowsA := len(matrixA)
	colsA := len(matrixA[0])
	rowsB := len(matrixB)
	colsB := len(matrixB[0])

	// Check if the matrices can be multiplied
	if colsA != rowsB {
		fmt.Println("Error: matrices cannot be multiplied")
		return
	}

	// Initialize the result matrix with the correct number of rows and columns
	result := make([][]int, rowsA)
	for i := range result {
		result[i] = make([]int, colsB)
	}

	// Perform matrix multiplication
	for i := 0; i < rowsA; i++ {
		for j := 0; j < colsB; j++ {
			sum := 0
			for k := 0; k < colsA; k++ {
				sum += matrixA[i][k] * matrixB[k][j]
			}
			result[i][j] = sum
		}
	}

	// Print the result matrix
	fmt.Println(result)
}
