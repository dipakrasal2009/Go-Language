package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	var choice int

	fmt.Println("Enter two numbers:")
	fmt.Scan(&num1, &num2)

	fmt.Println("\nArithmetic Operations:")
	fmt.Println("1. Addition")
	fmt.Println("2. Subtraction")
	fmt.Println("3. Multiplication")
	fmt.Println("4. Division")
	fmt.Println("5. Modulus")
	fmt.Print("Enter your choice (1-5): ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Printf("%.2f + %.2f = %.2f\n", num1, num2, num1+num2)
	case 2:
		fmt.Printf("%.2f - %.2f = %.2f\n", num1, num2, num1-num2)
	case 3:
		fmt.Printf("%.2f * %.2f = %.2f\n", num1, num2, num1*num2)
	case 4:
		if num2 == 0 {
			fmt.Println("Error: Division by zero")
		} else {
			fmt.Printf("%.2f / %.2f = %.2f\n", num1, num2, num1/num2)
		}
	case 5:
		fmt.Printf("%.0f %% %.0f = %.0f\n", num1, num2, float64(int(num1)%int(num2)))
	default:
		fmt.Println("Invalid choice")
	}
}
