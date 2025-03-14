package main

import (
	"./calculator"
	"fmt"
)

func main() {
	var num1, num2 float64
	var choice int

	fmt.Print("Enter first number: ")
	fmt.Scan(&num1)

	fmt.Print("Enter second number: ")
	fmt.Scan(&num2)

	fmt.Println("\nCalculator Operations:")
	fmt.Println("1. Addition")
	fmt.Println("2. Subtraction")
	fmt.Println("3. Multiplication")
	fmt.Println("4. Division")
	fmt.Print("Enter your choice (1-4): ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		result := calculator.Add(num1, num2)
		fmt.Printf("%.2f + %.2f = %.2f\n", num1, num2, result)
	case 2:
		result := calculator.Subtract(num1, num2)
		fmt.Printf("%.2f - %.2f = %.2f\n", num1, num2, result)
	case 3:
		result := calculator.Multiply(num1, num2)
		fmt.Printf("%.2f * %.2f = %.2f\n", num1, num2, result)
	case 4:
		if num2 == 0 {
			fmt.Println("Error: Division by zero")
		} else {
			result := calculator.Divide(num1, num2)
			fmt.Printf("%.2f / %.2f = %.2f\n", num1, num2, result)
		}
	default:
		fmt.Println("Invalid choice")
	}
}
