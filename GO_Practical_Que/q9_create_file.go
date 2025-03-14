package main

import (
	"fmt"
	"os"
)

func main() {
	var filename string
	fmt.Print("Enter the filename to create: ")
	fmt.Scan(&filename)

	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer file.Close()

	fmt.Println("File created successfully.")

	// Write content to the file
	fmt.Print("Enter content to write to the file (type 'EOF' on a new line to finish):\n")

	var line string
	for {
		fmt.Scan(&line)
		if line == "EOF" {
			break
		}
		_, err := file.WriteString(line + "\n")
		if err != nil {
			fmt.Printf("Error writing to file: %v\n", err)
			return
		}
	}

	fmt.Printf("Content written to %s successfully.\n", filename)
}
