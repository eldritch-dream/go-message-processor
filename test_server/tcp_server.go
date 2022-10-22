package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

// only needed below for sample processing

const message_str string = "{\"tail_number\": \"N20904\",\"engine_count\": 2,\"engine_name\": \"GEnx-1B\",\"latitude\": 39.11593389482025,\"longitude\": -67.32425341289998,\"altitude\": 36895.5,\"temperature\": -53.2}"

func main() {

	fmt.Println("Launching server...")

	// listen on all interfaces
	ln, _ := net.Listen("tcp", ":8081")

	// accept connection on port
	conn, _ := ln.Accept()

	// run loop forever (or until ctrl-c)
	for {
		newmessage := strings.ToUpper(message_str)
		// send new string back to client
		conn.Write([]byte(newmessage + "\n"))
		time.Sleep(30000)
	}
}
