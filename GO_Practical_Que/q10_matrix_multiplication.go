package main

import (
	"fmt"
)

func main() {
	var r1, c1, r2, c2 int

	// Input dimensions for first matrix
	fmt.Println("Enter dimensions of first matrix:")
	fmt.Print("Rows: ")
	fmt.Scan(&r1)
	fmt.Print("Columns: ")
	fmt.Scan(&c1)

	// Input dimensions for second matrix
	fmt.Println("Enter dimensions of second matrix:")
	fmt.Print("Rows: ")
	fmt.Scan(&r2)
	fmt.Print("Columns: ")
	fmt.Scan(&c2)

	// Check if multiplication is possible
	if c1 != r2 {
		fmt.Println("Matrix multiplication not possible. Number of columns in first matrix must equal number of rows in second matrix.")
		return
	}

	// Create matrices
	matrix1 := make([][]int, r1)
	for i := range matrix1 {
		matrix1[i] = make([]int, c1)
	}

	matrix2 := make([][]int, r2)
	for i := range matrix2 {
		matrix2[i] = make([]int, c2)
	}

	result := make([][]int, r1)
	for i := range result {
		result[i] = make([]int, c2)
	}

	// Input elements for first matrix
	fmt.Println("Enter elements for first matrix:")
	for i := 0; i < r1; i++ {
		for j := 0; j < c1; j++ {
			fmt.Printf("Enter element at position [%d][%d]: ", i, j)
			fmt.Scan(&matrix1[i][j])
		}
	}

	// Input elements for second matrix
	fmt.Println("Enter elements for second matrix:")
	for i := 0; i < r2; i++ {
		for j := 0; j < c2; j++ {
			fmt.Printf("Enter element at position [%d][%d]: ", i, j)
			fmt.Scan(&matrix2[i][j])
		}
	}

	// Perform matrix multiplication
	for i := 0; i < r1; i++ {
		for j := 0; j < c2; j++ {
			for k := 0; k < c1; k++ {
				result[i][j] += matrix1[i][k] * matrix2[k][j]
			}
		}
	}

	// Display the result
	fmt.Println("Resultant Matrix after multiplication:")
	for i := 0; i < r1; i++ {
		for j := 0; j < c2; j++ {
			fmt.Printf("%d\t", result[i][j])
		}
		fmt.Println()
	}
}
