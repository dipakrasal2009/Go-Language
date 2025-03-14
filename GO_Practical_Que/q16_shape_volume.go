package main

import (
	"fmt"
	"math"
)

// Shape interface with area and volume methods
type Shape interface {
	Area() float64
	Volume() float64
}

// Square type (actually a cube for 3D)
type Square struct {
	Side float64
}

// Rectangle type (actually a cuboid for 3D)
type Rectangle struct {
	Length float64
	Width  float64
	Height float64
}

// Implement Shape interface for Square
func (s Square) Area() float64 {
	return 6 * s.Side * s.Side // Surface area of cube
}

func (s Square) Volume() float64 {
	return math.Pow(s.Side, 3) // Volume of cube
}

// Implement Shape interface for Rectangle
func (r Rectangle) Area() float64 {
	return 2 * (r.Length*r.Width + r.Length*r.Height + r.Width*r.Height) // Surface area of cuboid
}

func (r Rectangle) Volume() float64 {
	return r.Length * r.Width * r.Height // Volume of cuboid
}

// Function to display shape details
func displayShape(s Shape) {
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Volume: %.2f\n", s.Volume())
}

func main() {
	// Create a square (cube)
	var side float64
	fmt.Print("Enter side length of square (cube): ")
	fmt.Scan(&side)
	square := Square{Side: side}

	// Create a rectangle (cuboid)
	var length, width, height float64
	fmt.Print("Enter length of rectangle (cuboid): ")
	fmt.Scan(&length)
	fmt.Print("Enter width of rectangle (cuboid): ")
	fmt.Scan(&width)
	fmt.Print("Enter height of rectangle (cuboid): ")
	fmt.Scan(&height)
	rectangle := Rectangle{Length: length, Width: width, Height: height}

	// Display square (cube) details
	fmt.Println("\nSquare (Cube) Details:")
	fmt.Printf("Side: %.2f\n", square.Side)
	displayShape(square)

	// Display rectangle (cuboid) details
	fmt.Println("\nRectangle (Cuboid) Details:")
	fmt.Printf("Length: %.2f\n", rectangle.Length)
	fmt.Printf("Width: %.2f\n", rectangle.Width)
	fmt.Printf("Height: %.2f\n", rectangle.Height)
	displayShape(rectangle)
}
