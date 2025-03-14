package main

import (
	"fmt"
)

func main() {
	// Create a buffered channel with capacity 5
	ch := make(chan int, 5)

	fmt.Printf("Initial Channel Capacity: %d\n", cap(ch))
	fmt.Printf("Initial Channel Length: %d\n", len(ch))

	// Store values in the channel
	ch <- 10
	ch <- 20
	ch <- 30

	fmt.Printf("\nAfter adding 3 values:\n")
	fmt.Printf("Channel Capacity: %d\n", cap(ch))
	fmt.Printf("Channel Length: %d\n", len(ch))

	// Read a value from the channel
	value1 := <-ch
	fmt.Printf("\nRead value: %d\n", value1)
	fmt.Printf("Channel Length after reading: %d\n", len(ch))

	// Read another value from the channel
	value2 := <-ch
	fmt.Printf("Read value: %d\n", value2)
	fmt.Printf("Channel Length after reading again: %d\n", len(ch))

	// Add more values
	ch <- 40
	ch <- 50
	ch <- 60

	fmt.Printf("\nFinal Channel Length: %d\n", len(ch))

	// Read all remaining values
	fmt.Println("\nRemaining values in the channel:")
	for len(ch) > 0 {
		value := <-ch
		fmt.Printf("%d ", value)
	}
	fmt.Println()
}
