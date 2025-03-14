package main

import (
	"fmt"
)

func main() {
	evenSum := 0
	oddSum := 0

	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			evenSum += i
		} else {
			oddSum += i
		}
	}

	fmt.Printf("Sum of even numbers from 1 to 100: %d\n", evenSum)
	fmt.Printf("Sum of odd numbers from 1 to 100: %d\n", oddSum)
}
