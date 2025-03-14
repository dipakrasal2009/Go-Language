package main

import (
	"fmt"
)

type Student struct {
	ID    int
	Name  string
	Age   int
	Grade string
}

// Method with pointer receiver
func (s *Student) show() {
	fmt.Println("\nStudent Details:")
	fmt.Println("----------------")
	fmt.Printf("ID: %d\n", s.ID)
	fmt.Printf("Name: %s\n", s.Name)
	fmt.Printf("Age: %d\n", s.Age)
	fmt.Printf("Grade: %s\n", s.Grade)
}

func main() {
	var student Student

	fmt.Println("Enter Student Details:")

	fmt.Print("ID: ")
	fmt.Scan(&student.ID)

	fmt.Print("Name: ")
	fmt.Scan(&student.Name)

	fmt.Print("Age: ")
	fmt.Scan(&student.Age)

	fmt.Print("Grade: ")
	fmt.Scan(&student.Grade)

	// Call the show method with a pointer receiver
	student.show()
}
