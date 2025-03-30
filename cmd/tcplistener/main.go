package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func getLinesChannel(f io.ReadCloser) <-chan string {
	stringChannel := make(chan string)

	go func() {
		defer f.Close()

		buffer := make([]byte, 1024) // Use a larger buffer for efficiency
		var currentLine []byte

		for {
			n, err := f.Read(buffer)
			if err != nil {
				if err == io.EOF && len(currentLine) > 0 {
					stringChannel <- string(currentLine)
				}
				break
			}

			// Add the new bytes to our current line
			currentLine = append(currentLine, buffer[:n]...)

			// Look for newlines and split the data accordingly
			for {
				i := 0
				for i < len(currentLine) {
					if currentLine[i] == '\n' {
						// Send the line and remove it from currentLine
						stringChannel <- string(currentLine[:i])
						currentLine = currentLine[i+1:]
						i = 0 // Start over with the new currentLine
					} else {
						i++
					}
				}
				break // No more newlines found
			}

			// If we have data but no newline, send it anyway
			if len(currentLine) > 0 {
				stringChannel <- string(currentLine)
				currentLine = currentLine[:0]
			}
		}

		close(stringChannel)
	}()

	return stringChannel
}

func main() {

	listener, err := net.Listen("tcp", ":42069")
	if err != nil {
		log.Fatal(err)
	}

	defer listener.Close()

	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Connection has been accepted")

		linesChannel := getLinesChannel(connection)
		for line := range linesChannel {
			fmt.Printf("%s\n", line)
		}

		fmt.Println("Connection has been closed!")
	}
}
