package main

import (
	"fmt"
)

func main() {
	var rows, cols int

	// Input matrix dimensions
	fmt.Print("Enter number of rows: ")
	fmt.Scan(&rows)
	fmt.Print("Enter number of columns: ")
	fmt.Scan(&cols)

	if rows <= 0 || cols <= 0 {
		fmt.Println("Invalid matrix dimensions")
		return
	}

	// Create and populate matrix
	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, cols)
	}

	fmt.Println("Enter matrix elements:")
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Printf("Enter element at position [%d][%d]: ", i, j)
			fmt.Scan(&matrix[i][j])
		}
	}

	// Create transpose matrix
	transpose := make([][]int, cols)
	for i := range transpose {
		transpose[i] = make([]int, rows)
	}

	// Calculate transpose
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			transpose[j][i] = matrix[i][j]
		}
	}

	// Display original matrix
	fmt.Println("\nOriginal Matrix:")
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Printf("%d\t", matrix[i][j])
		}
		fmt.Println()
	}

	// Display transpose matrix
	fmt.Println("\nTranspose Matrix:")
	for i := 0; i < cols; i++ {
		for j := 0; j < rows; j++ {
			fmt.Printf("%d\t", transpose[i][j])
		}
		fmt.Println()
	}
}
