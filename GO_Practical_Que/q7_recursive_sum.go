package main

import (
	"fmt"
)

// Function to calculate sum of digits recursively
func sumOfDigits(n int) int {
	// Base case
	if n == 0 {
		return 0
	}

	// Return rightmost digit plus sum of remaining digits
	return (n % 10) + sumOfDigits(n/10)
}

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	// Handle negative numbers
	if num < 0 {
		num = -num
	}

	sum := sumOfDigits(num)
	fmt.Printf("Sum of digits of %d is %d\n", num, sum)
}
