package main

import (
	"fmt"
)

// Method to copy array elements
func copyArray(source []int) []int {
	// Create a new array with the same size
	destination := make([]int, len(source))

	// Copy elements
	for i, val := range source {
		destination[i] = val
	}

	return destination
}

func main() {
	var n int
	fmt.Print("Enter the number of elements: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Invalid array size")
		return
	}

	// Create and populate source array
	sourceArray := make([]int, n)
	fmt.Println("Enter the elements:")
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &sourceArray[i])
	}

	// Copy array using the method
	destinationArray := copyArray(sourceArray)

	// Display both arrays
	fmt.Println("\nSource Array:")
	for _, val := range sourceArray {
		fmt.Printf("%d ", val)
	}

	fmt.Println("\nDestination Array (after copy):")
	for _, val := range destinationArray {
		fmt.Printf("%d ", val)
	}
	fmt.Println()
}
