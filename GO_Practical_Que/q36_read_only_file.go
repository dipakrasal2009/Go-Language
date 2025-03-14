package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	var filename string
	fmt.Print("Enter the filename to open in READ only mode: ")
	fmt.Scan(&filename)

	// Check if file exists
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		// Create a sample file if it doesn't exist
		sampleContent := "This is a sample file content.\nIt is created because the specified file did not exist.\nYou can now read this file in READ only mode."
		err := ioutil.WriteFile(filename, []byte(sampleContent), 0644)
		if err != nil {
			fmt.Printf("Error creating sample file: %v\n", err)
			return
		}
		fmt.Printf("Created sample file '%s' as it did not exist.\n\n", filename)
	}

	// Open file in READ only mode
	file, err := os.OpenFile(filename, os.O_RDONLY, 0)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	fmt.Printf("File '%s' opened successfully in READ only mode.\n", filename)

	// Read file content
	content, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Display file content
	fmt.Println("\nFile Content:")
	fmt.Println("-------------")
	fmt.Println(string(content))

	// Try to write to the file (should fail as it's opened in READ only mode)
	_, err = file.WriteString("This should not be written as the file is in READ only mode.")
	if err != nil {
		fmt.Printf("\nAttempt to write to the READ only file failed as expected with error: %v\n", err)
	} else {
		fmt.Println("\nUnexpected: Was able to write to READ only file.")
	}
}
