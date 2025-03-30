package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	udpAddres, err := net.ResolveUDPAddr("udp", "localhost:42069")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	udpConnection, err := net.DialUDP("udp", nil, udpAddres)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer udpConnection.Close()

	fmt.Printf("Sending to %s. Type your message and press Enter to send. Press Ctrl+C to exit.\n", "locahost:42069")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">")
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		_, err = udpConnection.Write([]byte(line))
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}
}
