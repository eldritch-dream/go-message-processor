package main

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"github.com/eldritch-dream/go-message-processor/m/v2/flightMessage"
)

// const server_addr string = "data.salad.com:5000"
const server_addr string = "localhost:8081"

func main() {
	fmt.Println("Init Processor")

	tcpAddr, err := net.ResolveTCPAddr("tcp", server_addr)
	if err != nil {
		fmt.Printf("ResolveTCPAddr failed with error: %s", err)
		os.Exit(1)
	}

	fmt.Printf("TCP Address resolved to: %s", tcpAddr.String())

	conn, err := net.Dial("tcp", server_addr)
	if err != nil {
		fmt.Printf("Dial tcp failed with error: %s", err)
		os.Exit(1)
	}

	connReader := bufio.NewReader(conn)
	buf := make([]byte, 1024)

	for {
		_, err := connReader.Read(buf)
		if err != nil {
			fmt.Printf("Error reading from connection: %s", err)
			os.Exit(1)
		}

		decodedMessage := flightMessage.CreateMessageFromBinary(buf)

		fmt.Println("Decoded message was: ", decodedMessage)

	}

}
