package main

import (
	"fmt"
	"strconv"
)

// Function to calculate sum of squares
func calculateSquares(number int, squareCh chan<- int) {
	sum := 0
	// Convert number to string to iterate through digits
	numStr := strconv.Itoa(number)

	for _, digitChar := range numStr {
		digit, _ := strconv.Atoi(string(digitChar))
		sum += digit * digit
	}

	squareCh <- sum
}

// Function to calculate sum of cubes
func calculateCubes(number int, cubeCh chan<- int) {
	sum := 0
	// Convert number to string to iterate through digits
	numStr := strconv.Itoa(number)

	for _, digitChar := range numStr {
		digit, _ := strconv.Atoi(string(digitChar))
		sum += digit * digit * digit
	}

	cubeCh <- sum
}

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	// Create channels
	squareCh := make(chan int)
	cubeCh := make(chan int)

	// Start goroutines
	go calculateSquares(num, squareCh)
	go calculateCubes(num, cubeCh)

	// Receive results
	squareSum := <-squareCh
	cubeSum := <-cubeCh

	fmt.Printf("Number: %d\n", num)
	fmt.Printf("Sum of squares of digits: %d\n", squareSum)
	fmt.Printf("Sum of cubes of digits: %d\n", cubeSum)
}
