package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	var filename string
	fmt.Print("Enter the filename to get information: ")
	fmt.Scan(&filename)

	fileInfo, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("File '%s' does not exist\n", filename)
		} else {
			fmt.Printf("Error: %v\n", err)
		}
		return
	}

	fmt.Println("\nFile Information:")
	fmt.Println("----------------")
	fmt.Printf("Name: %s\n", fileInfo.Name())
	fmt.Printf("Size: %d bytes\n", fileInfo.Size())
	fmt.Printf("Mode: %v\n", fileInfo.Mode())
	fmt.Printf("Last Modified: %v\n", fileInfo.ModTime().Format(time.RFC1123))
	fmt.Printf("Is Directory: %t\n", fileInfo.IsDir())
}
