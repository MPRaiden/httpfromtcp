package main

import (
	"fmt"
	"log"
	"net"

	"github.com/MPRaiden/httpfromtcp/internal/request"
)

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
		defer connection.Close()

		fmt.Println("Connection has been accepted")

		req, err := request.RequestFromReader(connection)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Request line:")
		fmt.Printf("- Method: %v\n", req.RequestLine.Method)
		fmt.Printf("- Target: %v\n", req.RequestLine.RequestTarget)
		fmt.Printf("- Version: %v\n", req.RequestLine.HttpVersion)
	}
}
