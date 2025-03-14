package main

import (
	"fmt"
)

// Interface with a method
type Vehicle interface {
	Details() string
}

// Car type
type Car struct {
	Brand  string
	Model  string
	Year   int
	Engine string
}

// Bike type
type Bike struct {
	Brand string
	Model string
	Year  int
	Type  string
}

// Implement Vehicle interface for Car
func (c Car) Details() string {
	return fmt.Sprintf("Car: %s %s (%d) with %s engine", c.Brand, c.Model, c.Year, c.Engine)
}

// Implement Vehicle interface for Bike
func (b Bike) Details() string {
	return fmt.Sprintf("Bike: %s %s (%d) - %s type", b.Brand, b.Model, b.Year, b.Type)
}

func main() {
	// Create vehicles slice containing different vehicle types
	vehicles := []Vehicle{
		Car{Brand: "Toyota", Model: "Corolla", Year: 2020, Engine: "1.8L"},
		Bike{Brand: "Honda", Model: "CBR", Year: 2021, Type: "Sports"},
		Car{Brand: "Tesla", Model: "Model 3", Year: 2022, Engine: "Electric"},
	}

	// Display vehicle details and use type assertion
	fmt.Println("Vehicle Details:")
	for i, v := range vehicles {
		fmt.Printf("\nVehicle %d: %s\n", i+1, v.Details())

		// Type assertion to access specific fields
		if car, ok := v.(Car); ok {
			fmt.Printf("This is a car with %s engine\n", car.Engine)
		}

		if bike, ok := v.(Bike); ok {
			fmt.Printf("This is a %s bike\n", bike.Type)
		}
	}
}
