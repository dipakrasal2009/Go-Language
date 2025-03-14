package main

import (
	"fmt"
)

// Function that returns multiple values (add, subtract, multiply, divide)
func calculate(a, b float64) (float64, float64, float64, float64) {
	add := a + b
	subtract := a - b
	multiply := a * b

	var divide float64
	if b != 0 {
		divide = a / b
	} else {
		divide = 0 // Avoid division by zero
	}

	return add, subtract, multiply, divide
}

func main() {
	var num1, num2 float64

	fmt.Print("Enter first number: ")
	fmt.Scan(&num1)

	fmt.Print("Enter second number: ")
	fmt.Scan(&num2)

	// Call function that returns multiple values
	sum, difference, product, quotient := calculate(num1, num2)

	fmt.Printf("\nResults for %.2f and %.2f:\n", num1, num2)
	fmt.Printf("Addition: %.2f + %.2f = %.2f\n", num1, num2, sum)
	fmt.Printf("Subtraction: %.2f - %.2f = %.2f\n", num1, num2, difference)
	fmt.Printf("Multiplication: %.2f * %.2f = %.2f\n", num1, num2, product)

	if num2 != 0 {
		fmt.Printf("Division: %.2f / %.2f = %.2f\n", num1, num2, quotient)
	} else {
		fmt.Println("Division: Cannot divide by zero")
	}
}
