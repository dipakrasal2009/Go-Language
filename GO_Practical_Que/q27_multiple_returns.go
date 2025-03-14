package main

import (
	"fmt"
)

// Function that returns multiple values
func calculate(a, b int) (sum, difference, product, quotient float64) {
	sum = float64(a + b)
	difference = float64(a - b)
	product = float64(a * b)

	if b != 0 {
		quotient = float64(a) / float64(b)
	} else {
		quotient = 0 // Avoid division by zero
	}

	return
}

func main() {
	var num1, num2 int

	fmt.Print("Enter first number: ")
	fmt.Scan(&num1)

	fmt.Print("Enter second number: ")
	fmt.Scan(&num2)

	sum, difference, product, quotient := calculate(num1, num2)

	fmt.Printf("\nResults for %d and %d:\n", num1, num2)
	fmt.Printf("Sum: %.2f\n", sum)
	fmt.Printf("Difference: %.2f\n", difference)
	fmt.Printf("Product: %.2f\n", product)

	if num2 != 0 {
		fmt.Printf("Quotient: %.2f\n", quotient)
	} else {
		fmt.Println("Quotient: Division by zero not possible")
	}
}
