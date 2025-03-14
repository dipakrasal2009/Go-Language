package main

import (
	"fmt"
)

// Function to generate Fibonacci series and send through channel
func fibonacci(n int, ch chan<- int) {
	a, b := 0, 1

	// Send first two terms
	ch <- a
	if n > 1 {
		ch <- b
	}

	// Generate and send remaining terms
	for i := 2; i < n; i++ {
		next := a + b
		ch <- next
		a, b = b, next
	}

	close(ch)
}

func main() {
	var n int
	fmt.Print("Enter the number of terms for Fibonacci series: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Invalid input. Number of terms should be positive.")
		return
	}

	// Create channel
	ch := make(chan int, n)

	// Start goroutine to generate Fibonacci series
	go fibonacci(n, ch)

	// Read and display Fibonacci series from channel
	fmt.Printf("Fibonacci Series of %d terms: ", n)
	for num := range ch {
		fmt.Printf("%d ", num)
	}
	fmt.Println()
}
