package main

import (
	"./rectangle"
	"fmt"
)

func main() {
	var length, width float64

	fmt.Print("Enter the length of the rectangle: ")
	fmt.Scan(&length)

	fmt.Print("Enter the width of the rectangle: ")
	fmt.Scan(&width)

	area := rectangle.Calculate(length, width)

	fmt.Printf("Area of rectangle with length %.2f and width %.2f is: %.2f\n",
		length, width, area)
}
