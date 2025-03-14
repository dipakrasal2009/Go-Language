package main

import (
	"fmt"
	"math"
)

// Shape interface with area and perimeter methods
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Circle type
type Circle struct {
	Radius float64
}

// Rectangle type
type Rectangle struct {
	Length float64
	Width  float64
}

// Implement Shape interface for Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Implement Shape interface for Rectangle
func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Width)
}

// Function to display shape details
func displayShape(s Shape) {
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
}

func main() {
	// Create a circle
	var radius float64
	fmt.Print("Enter radius of circle: ")
	fmt.Scan(&radius)
	circle := Circle{Radius: radius}

	// Create a rectangle
	var length, width float64
	fmt.Print("Enter length of rectangle: ")
	fmt.Scan(&length)
	fmt.Print("Enter width of rectangle: ")
	fmt.Scan(&width)
	rectangle := Rectangle{Length: length, Width: width}

	// Display circle details
	fmt.Println("\nCircle Details:")
	fmt.Printf("Radius: %.2f\n", circle.Radius)
	displayShape(circle)

	// Display rectangle details
	fmt.Println("\nRectangle Details:")
	fmt.Printf("Length: %.2f\n", rectangle.Length)
	fmt.Printf("Width: %.2f\n", rectangle.Width)
	displayShape(rectangle)
}
