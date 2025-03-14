package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Enter the number of terms for Fibonacci series: ")
	fmt.Scan(&n)

	fmt.Printf("Fibonacci Series of %d terms: ", n)

	// Handle special cases
	if n <= 0 {
		fmt.Println("Invalid input. Number of terms should be positive.")
		return
	}

	if n == 1 {
		fmt.Println("0")
		return
	}

	// Print the first two terms
	fmt.Print("0 1")

	a, b := 0, 1

	// Generate and print remaining terms
	for i := 2; i < n; i++ {
		next := a + b
		fmt.Printf(" %d", next)
		a, b = b, next
	}
	fmt.Println()
}
