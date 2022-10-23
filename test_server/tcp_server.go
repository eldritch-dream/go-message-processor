package main

import (
	"encoding/binary"
	"fmt"
	"math"
	"net"
	"time"
)

// only needed below for sample processing

const HEADER string = "AIR"
const TAIL_NUMBER_SIZE uint32 = 10
const ENGINE_COUNT uint32 = 4
const ENGINE_NAME_SIZE uint32 = 7
const LATITUDE float64 = 39.11593389482025
const LONGITUDE float64 = -67.32425341289998
const ALTITUDE float64 = 36895.5
const TEMPERATURE float64 = -53.2

const TAIL_NUMBER string = "1234ABCDEF"
const ENGINE_NAME string = "GEnx-1B"

func main() {

	fmt.Println("Launching server...")

	// listen on all interfaces
	ln, _ := net.Listen("tcp", ":8081")

	// accept connection on port
	conn, _ := ln.Accept()

	// run loop forever (or until ctrl-c)
	for {
		// send new string back to client
		var messageBytes []byte
		messageBytes = append(messageBytes, HEADER...)

		messageBytes = append(messageBytes, Uint32bytes(TAIL_NUMBER_SIZE)...)
		messageBytes = append(messageBytes, TAIL_NUMBER...)

		messageBytes = append(messageBytes, Uint32bytes(ENGINE_COUNT)...)
		messageBytes = append(messageBytes, Uint32bytes(ENGINE_NAME_SIZE)...)
		messageBytes = append(messageBytes, ENGINE_NAME...)

		messageBytes = append(messageBytes, Float64bytes(LATITUDE)...)
		messageBytes = append(messageBytes, Float64bytes(LONGITUDE)...)

		messageBytes = append(messageBytes, Float64bytes(ALTITUDE)...)

		messageBytes = append(messageBytes, Float64bytes(TEMPERATURE)...)

		conn.Write(messageBytes)
		time.Sleep(time.Second * 5)
	}
}

func Uint32bytes(num uint32) []byte {
	bytes := make([]byte, 4)
	binary.BigEndian.PutUint32(bytes, num)
	return bytes
}

func Float64bytes(float float64) []byte {
	bits := math.Float64bits(float)
	bytes := make([]byte, 8)
	binary.BigEndian.PutUint64(bytes, bits)
	return bytes
}
