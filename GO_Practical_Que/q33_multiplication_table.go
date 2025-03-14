package main

import (
	"fmt"
)

// Function to print multiplication table
func printMultiplicationTable(num, rows int) {
	fmt.Printf("Multiplication Table for %d:\n", num)
	fmt.Println("---------------------------")

	for i := 1; i <= rows; i++ {
		fmt.Printf("%d × %d = %d\n", num, i, num*i)
	}
}

func main() {
	var num, rows int

	fmt.Print("Enter a number for multiplication table: ")
	fmt.Scan(&num)

	fmt.Print("Enter the number of rows: ")
	fmt.Scan(&rows)

	if rows <= 0 {
		fmt.Println("Number of rows should be positive")
		return
	}

	printMultiplicationTable(num, rows)
}
