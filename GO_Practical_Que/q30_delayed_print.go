package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Printing numbers from 0 to 10 with random delays:")

	for i := 0; i <= 10; i++ {
		// Generate random delay between 0 and 250ms
		delay := time.Duration(rand.Intn(251)) * time.Millisecond

		// Print number with delay information
		fmt.Printf("Number %d (delay: %v)\n", i, delay)

		// Wait for the specified delay
		time.Sleep(delay)
	}

	fmt.Println("Done!")
}
