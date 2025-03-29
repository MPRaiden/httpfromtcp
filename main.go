package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("messages.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	buffer := make([]byte, 8) // Read 8 bytes at a time

	for err != io.EOF {
		_, err = file.Read(buffer)
		fmt.Printf("read: %s\n", string(buffer))
	}
}
