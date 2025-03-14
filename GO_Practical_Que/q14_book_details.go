package main

import (
	"fmt"
)

type Book struct {
	BookID int
	Title  string
	Author string
	Price  float64
}

func main() {
	var n int
	fmt.Print("Enter the number of books: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Invalid number of books")
		return
	}

	books := make([]Book, n)

	// Input book details
	for i := 0; i < n; i++ {
		fmt.Printf("\nEnter details for Book %d:\n", i+1)

		fmt.Print("Book ID: ")
		fmt.Scan(&books[i].BookID)

		fmt.Print("Title: ")
		fmt.Scan(&books[i].Title)

		fmt.Print("Author: ")
		fmt.Scan(&books[i].Author)

		fmt.Print("Price: ")
		fmt.Scan(&books[i].Price)
	}

	// Display book details
	fmt.Println("\nBook Details:")
	fmt.Println("------------")
	for i, book := range books {
		fmt.Printf("\nBook %d:\n", i+1)
		fmt.Printf("Book ID: %d\n", book.BookID)
		fmt.Printf("Title: %s\n", book.Title)
		fmt.Printf("Author: %s\n", book.Author)
		fmt.Printf("Price: $%.2f\n", book.Price)
	}
}
