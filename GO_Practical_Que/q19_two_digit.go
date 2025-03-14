package main

import (
	"fmt"
)

// Function to check if a number is two-digit
func isTwoDigit(num int) bool {
	return num >= 10 && num <= 99
}

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	// Check for negative numbers
	if num < 0 {
		num = -num
	}

	if isTwoDigit(num) {
		fmt.Printf("%d is a two-digit number\n", num)
	} else {
		fmt.Printf("%d is not a two-digit number\n", num)
	}
}
