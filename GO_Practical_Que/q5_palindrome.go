package main

import (
	"fmt"
)

// Function to check if a number is palindrome
func isPalindrome(num int) bool {
	original := num
	reversed := 0

	for num > 0 {
		remainder := num % 10
		reversed = reversed*10 + remainder
		num /= 10
	}

	return original == reversed
}

func main() {
	var num int
	fmt.Print("Enter a number to check if it's a palindrome: ")
	fmt.Scan(&num)

	if isPalindrome(num) {
		fmt.Printf("%d is a palindrome number\n", num)
	} else {
		fmt.Printf("%d is not a palindrome number\n", num)
	}
}
