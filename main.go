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

		buffer := make([]byte, 8)
		var currentLine []byte

		for {
			n, err := f.Read(buffer)
			if err != nil {
				if err == io.EOF {
					// End of file reached
					if len(currentLine) > 0 {
						stringChannel <- string(currentLine)
					}
					break
				}
				break
			}

			// Process the bytes read
			for i := range buffer[:n] {
				if buffer[i] == '\n' {
					// Complete line found, send it to the channel
					stringChannel <- string(currentLine)
					currentLine = currentLine[:0] // Clear the slice
				} else {
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
