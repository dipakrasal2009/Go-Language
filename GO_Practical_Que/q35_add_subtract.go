package main

import (
	"fmt"
)

// Function that returns multiple values (add, subtract)
func calculate(a, b int) (int, int) {
	add := a + b
	subtract := a - b
	return add, subtract
}

func main() {
	var num1, num2 int

	fmt.Print("Enter first number: ")
	fmt.Scan(&num1)

	fmt.Print("Enter second number: ")
	fmt.Scan(&num2)

	// Call function that returns multiple values
	sum, difference := calculate(num1, num2)

	fmt.Printf("\nResults for %d and %d:\n", num1, num2)
	fmt.Printf("Addition: %d + %d = %d\n", num1, num2, sum)
	fmt.Printf("Subtraction: %d - %d = %d\n", num1, num2, difference)
}
