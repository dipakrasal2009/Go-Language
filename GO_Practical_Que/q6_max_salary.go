package main

import (
	"fmt"
)

type Employee struct {
	EmpNo  int
	Name   string
	Salary float64
}

func main() {
	var n int
	fmt.Print("Enter the number of employees: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Invalid number of employees")
		return
	}

	employees := make([]Employee, n)

	// Input employee details
	for i := 0; i < n; i++ {
		fmt.Printf("\nEnter details for employee %d:\n", i+1)

		fmt.Print("Employee Number: ")
		fmt.Scan(&employees[i].EmpNo)

		fmt.Print("Employee Name: ")
		fmt.Scan(&employees[i].Name)

		fmt.Print("Salary: ")
		fmt.Scan(&employees[i].Salary)
	}

	// Find employee with maximum salary
	maxSalaryIdx := 0
	for i := 1; i < n; i++ {
		if employees[i].Salary > employees[maxSalaryIdx].Salary {
			maxSalaryIdx = i
		}
	}

	// Display employee with maximum salary
	fmt.Println("\nEmployee with Maximum Salary:")
	fmt.Println("-----------------------------")
	fmt.Printf("Employee Number: %d\n", employees[maxSalaryIdx].EmpNo)
	fmt.Printf("Name: %s\n", employees[maxSalaryIdx].Name)
	fmt.Printf("Salary: %.2f\n", employees[maxSalaryIdx].Salary)
}
