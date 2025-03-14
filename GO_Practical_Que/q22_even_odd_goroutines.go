package main

import (
	"fmt"
	"sync"
)

// Function to process even numbers
func processEven(evenCh <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Even numbers:")
	for num := range evenCh {
		fmt.Printf("%d ", num)
	}
	fmt.Println()
}

// Function to process odd numbers
func processOdd(oddCh <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Odd numbers:")
	for num := range oddCh {
		fmt.Printf("%d ", num)
	}
	fmt.Println()
}

func main() {
	var n int
	fmt.Print("Enter the number of elements: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Invalid slice size")
		return
	}

	// Create and populate the slice
	numbers := make([]int, n)
	fmt.Println("Enter the elements:")
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &numbers[i])
	}

	// Create channels for even and odd numbers
	evenCh := make(chan int)
	oddCh := make(chan int)

	// Create wait group for synchronization
	var wg sync.WaitGroup
	wg.Add(2)

	// Start goroutines to process even and odd numbers
	go processEven(evenCh, &wg)
	go processOdd(oddCh, &wg)

	// Check numbers and send to appropriate channels
	for _, num := range numbers {
		if num%2 == 0 {
			evenCh <- num
		} else {
			oddCh <- num
		}
	}

	// Close channels to signal completion
	close(evenCh)
	close(oddCh)

	// Wait for goroutines to finish
	wg.Wait()
}
