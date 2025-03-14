package main

import (
	"fmt"
)

func main() {
	// Create a channel
	ch := make(chan int)

	// Run a goroutine to send values to the channel
	go func() {
		fmt.Println("Sending values to channel...")
		for i := 1; i <= 5; i++ {
			ch <- i
			fmt.Printf("Sent: %d\n", i)
		}
		// Close the channel when done sending
		close(ch)
		fmt.Println("Channel closed")
	}()

	// Read values from the channel using a for-range loop
	fmt.Println("Reading values from channel:")
	for value := range ch {
		fmt.Printf("Received: %d\n", value)
	}

	fmt.Println("Channel has been closed and all values have been received")
}
