package main

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"github.com/eldritch-dream/go-message-processor/m/v2/flightMessage"
)

// const server_addr string = "data.salad.com:5000"
const server_addr string = "127.0.0.1:8082"

func main() {
	fmt.Println("Init Processor")

	tcpAddr, err := net.ResolveTCPAddr("tcp", server_addr)
	if err != nil {
		fmt.Printf("ResolveTCPAddr failed with error: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("TCP Address resolved to: %s\n", tcpAddr.String())

	for {

		conn, err := net.Dial("tcp", server_addr)
		if err != nil {
			fmt.Printf("Dial tcp failed with error: %s\n", err)
			continue
		}

		defer conn.Close()

		connReader := bufio.NewReader(conn)
		//This buffer is probably overkill but I don't know the max size
		buf := make([]byte, 1024)

		numBytesRead, err := connReader.Read(buf)
		if err != nil {
			fmt.Printf("Error reading from connection: %s\n", err)
			conn.Close()
			continue
		}

		if numBytesRead < 33 {
			fmt.Println("Read less than 33 bytes from connection, message is probably incomplete")
			continue
		}

		// fmt.Println("Read Byte buffer with content: ", buf)

		decodedMessage, err := flightMessage.CreateMessageFromBinary(buf)
		if err != nil {
			fmt.Printf("Error decoding message from bytes: %s\n", err)
			conn.Close()
			continue
		}

		fmt.Printf("%+v\n", decodedMessage)

	}

}
