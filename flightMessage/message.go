package flightMessage

import (
	"bytes"
	"encoding/binary"
)

type FlightMessage struct {
	tail_number  string  //The international aircraft registration. A unique code assigned to the aircraft.
	engine_count int     //The number of engines on the aircraft.
	engine_name  string  //The engine name.
	latitude     float64 //The latitude in degrees.
	longitude    float64 //The longitude in degrees.
	altitude     float64 //The altitude in feet.
	temperature  float64 //The temperature in degrees Fahrenheit.
}

var HEADER = []byte{0x41, 0x49, 0x52}

func CreateMessageFromBinary(messageBytes []byte) *FlightMessage {

	headerBytes := messageBytes[0:3]

	if bytes.Compare(headerBytes, HEADER) == 0 {

		messageReader := bytes.NewReader(messageBytes[3:])

		flightMessage := &FlightMessage{}
		binary.Read(messageReader, binary.BigEndian, flightMessage)
		return flightMessage
	}

	return nil

}
