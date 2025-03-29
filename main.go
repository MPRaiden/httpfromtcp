package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func getLinesChannel(f io.ReadCloser) <-chan string {
	stringChannel := make(chan string)

	go func() {
		defer f.Close()

		buffer := make([]byte, 8) // Read 8 bytes at a time
		var currentLine []byte    // Use byte slice instead of string

		for {
			n, err := f.Read(buffer)
			if err != nil {
				if err == io.EOF {
					// End of file reached
					if len(currentLine) > 0 {
						// Send the last line if it exists
						stringChannel <- string(currentLine)
					}
					break
				}
				// Handle other errors if needed
				break
			}

			// Process the bytes read
			for i := range buffer[:n] {
				if buffer[i] == '\n' {
					// Complete line found, send it to the channel
					stringChannel <- string(currentLine)
					currentLine = currentLine[:0] // Clear the slice
				} else {
					// Add byte to current line
					currentLine = append(currentLine, buffer[i])
				}
			}
		}

		close(stringChannel)
	}()

	return stringChannel
}

func main() {
	file, err := os.Open("messages.txt")
	if err != nil {
		log.Fatal(err)
	}

	linesChannel := getLinesChannel(file)

	for line := range linesChannel {
		fmt.Printf("read: %s\n", line)
	}

}
