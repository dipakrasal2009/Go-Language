package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Enter the number of elements: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Invalid array size")
		return
	}

	arr := make([]int, n)

	fmt.Println("Enter the elements:")
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &arr[i])
	}

	// Bubble sort algorithm
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// Swap elements
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	fmt.Println("Sorted array in ascending order:")
	for _, val := range arr {
		fmt.Printf("%d ", val)
	}
	fmt.Println()
}
