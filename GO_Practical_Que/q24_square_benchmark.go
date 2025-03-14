package main

import (
	"fmt"
	"testing"
)

// Function to find square of a number
func Square(n int) int {
	return n * n
}

// Benchmark function
func BenchmarkSquare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Square(7)
	}
}

func main() {
	var num int
	fmt.Print("Enter a number to find its square: ")
	fmt.Scan(&num)

	result := Square(num)
	fmt.Printf("Square of %d is %d\n", num, result)

	fmt.Println("\nTo run the benchmark, save this code as square_test.go and execute:")
	fmt.Println("go test -bench=.")
}
