package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

// Book structure for XML parsing
type Book struct {
	XMLName  xml.Name `xml:"book"`
	ID       int      `xml:"id,attr"`
	Title    string   `xml:"title"`
	Author   string   `xml:"author"`
	Year     int      `xml:"year"`
	Price    float64  `xml:"price"`
	Category string   `xml:"category"`
}

// Bookstore structure
type Bookstore struct {
	XMLName xml.Name `xml:"bookstore"`
	Books   []Book   `xml:"book"`
}

func main() {
	// Create sample XML file if it doesn't exist
	xmlFile := "bookstore.xml"
	if _, err := os.Stat(xmlFile); os.IsNotExist(err) {
		sampleXML := `<?xml version="1.0" encoding="UTF-8"?>
<bookstore>
  <book id="1">
    <title>Harry Potter</title>
    <author>J.K. Rowling</author>
    <year>2005</year>
    <price>29.99</price>
    <category>Fantasy</category>
  </book>
  <book id="2">
    <title>The Go Programming Language</title>
    <author>Alan Donovan</author>
    <year>2015</year>
    <price>39.99</price>
    <category>Programming</category>
  </book>
  <book id="3">
    <title>The Alchemist</title>
    <author>Paulo Coelho</author>
    <year>1988</year>
    <price>19.99</price>
    <category>Fiction</category>
  </book>
</bookstore>`

		err := ioutil.WriteFile(xmlFile, []byte(sampleXML), 0644)
		if err != nil {
			fmt.Printf("Error creating sample XML file: %v\n", err)
			return
		}
		fmt.Printf("Created sample XML file: %s\n\n", xmlFile)
	}

	// Read XML file
	xmlData, err := ioutil.ReadFile(xmlFile)
	if err != nil {
		fmt.Printf("Error reading XML file: %v\n", err)
		return
	}

	// Unmarshal XML into structure
	var bookstore Bookstore
	err = xml.Unmarshal(xmlData, &bookstore)
	if err != nil {
		fmt.Printf("Error unmarshalling XML: %v\n", err)
		return
	}

	// Display structure contents
	fmt.Println("Bookstore Contents:")
	fmt.Println("------------------")
	for _, book := range bookstore.Books {
		fmt.Printf("Book ID: %d\n", book.ID)
		fmt.Printf("Title: %s\n", book.Title)
		fmt.Printf("Author: %s\n", book.Author)
		fmt.Printf("Year: %d\n", book.Year)
		fmt.Printf("Price: $%.2f\n", book.Price)
		fmt.Printf("Category: %s\n", book.Category)
		fmt.Println()
	}
}
