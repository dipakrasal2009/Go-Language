package main

import (
	"fmt"
	"os"
)

func main() {
	var filename string
	fmt.Print("Enter the filename to append content: ")
	fmt.Scan(&filename)

	// Open file in append mode
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	fmt.Printf("File '%s' opened successfully for appending.\n", filename)

	// Get content to append
	fmt.Println("Enter content to append (type 'EOF' on a new line to finish):")

	var line string
	for {
		fmt.Scan(&line)
		if line == "EOF" {
			break
		}

		// Append to file
		_, err := file.WriteString(line + "\n")
		if err != nil {
			fmt.Printf("Error writing to file: %v\n", err)
			return
		}
	}

	fmt.Printf("Content appended to '%s' successfully.\n", filename)
}
