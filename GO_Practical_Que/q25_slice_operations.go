package main

import (
	"fmt"
)

func main() {
	// Create an original slice
	original := []int{1, 2, 3, 4, 5}
	fmt.Println("Original slice:", original)

	// Append elements
	original = append(original, 6, 7)
	fmt.Println("After appending 6, 7:", original)

	// Create another slice
	another := []int{8, 9, 10}

	// Append another slice
	original = append(original, another...)
	fmt.Println("After appending another slice:", original)

	// Create a slice with specific capacity
	withCapacity := make([]int, 3, 10)
	fmt.Printf("Slice with length %d and capacity %d: %v\n", len(withCapacity), cap(withCapacity), withCapacity)

	// Copy elements
	destination := make([]int, 5)
	copied := copy(destination, original)
	fmt.Printf("Copied %d elements: %v\n", copied, destination)

	// Slicing
	sliced := original[2:5]
	fmt.Println("Sliced (index 2 to 4):", sliced)

	// Remove an element at index 3
	index := 3
	original = append(original[:index], original[index+1:]...)
	fmt.Printf("After removing element at index %d: %v\n", index, original)

	// Clear a slice
	original = original[:0]
	fmt.Println("After clearing:", original)
}
