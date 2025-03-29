package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("messages.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	buffer := make([]byte, 8) // Read 8 bytes at a time
	var currentLine string

	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		data := string(buffer[:n]) // Only consider bytes that were actually read
		parts := strings.Split(data, "\n")

		// Process all parts except the last one
		for _, part := range parts[:len(parts)-1] {
			currentLine += part
			fmt.Printf("read: %s\n", currentLine)
			currentLine = ""
		}

		// Add the last part to currentLine
		currentLine += parts[len(parts)-1]
	}

	// Don't forget to print the last line if it's not empty
	if currentLine != "" {
		fmt.Printf("read: %s\n", currentLine)
	}
}

