package main

import (
	"fmt"
)

type Student struct {
	RollNo int
	Name   string
	Mark1  float64
	Mark2  float64
	Mark3  float64
	Total  float64
	Avg    float64
}

func main() {
	var n int
	fmt.Print("Enter the number of students: ")
	fmt.Scan(&n)

	students := make([]Student, n)

	for i := 0; i < n; i++ {
		fmt.Printf("\nEnter details for student %d:\n", i+1)

		fmt.Print("Roll No: ")
		fmt.Scan(&students[i].RollNo)

		fmt.Print("Name: ")
		fmt.Scan(&students[i].Name)

		fmt.Print("Mark 1: ")
		fmt.Scan(&students[i].Mark1)

		fmt.Print("Mark 2: ")
		fmt.Scan(&students[i].Mark2)

		fmt.Print("Mark 3: ")
		fmt.Scan(&students[i].Mark3)

		// Calculate total and average
		students[i].Total = students[i].Mark1 + students[i].Mark2 + students[i].Mark3
		students[i].Avg = students[i].Total / 3.0
	}

	// Display student details
	fmt.Println("\nStudent Details:")
	fmt.Println("----------------")
	for i, student := range students {
		fmt.Printf("\nStudent %d:\n", i+1)
		fmt.Printf("Roll No: %d\n", student.RollNo)
		fmt.Printf("Name: %s\n", student.Name)
		fmt.Printf("Mark 1: %.2f\n", student.Mark1)
		fmt.Printf("Mark 2: %.2f\n", student.Mark2)
		fmt.Printf("Mark 3: %.2f\n", student.Mark3)
		fmt.Printf("Total: %.2f\n", student.Total)
		fmt.Printf("Average: %.2f\n", student.Avg)
	}
}
