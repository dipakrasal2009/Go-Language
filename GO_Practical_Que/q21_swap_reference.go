package main

import (
	"fmt"
)

// Function to swap two numbers using pointers (call by reference)
func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	var num1, num2 int

	fmt.Print("Enter first number: ")
	fmt.Scan(&num1)

	fmt.Print("Enter second number: ")
	fmt.Scan(&num2)

	fmt.Printf("\nBefore swapping:\n")
	fmt.Printf("First number: %d\n", num1)
	fmt.Printf("Second number: %d\n", num2)

	// Call swap function with addresses of variables
	swap(&num1, &num2)

	fmt.Printf("\nAfter swapping:\n")
	fmt.Printf("First number: %d\n", num1)
	fmt.Printf("Second number: %d\n", num2)
}
