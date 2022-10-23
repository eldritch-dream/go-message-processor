package flightMessage

import (
	"bytes"
	"encoding/binary"
	"fmt"
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

	headerBytes := make([]byte, 3)
	messageReader := bytes.NewReader(messageBytes)
	messageReader.Read(headerBytes)
	fmt.Println("Header bytes are: ", headerBytes)

	if bytes.Compare(headerBytes, HEADER) == 0 {

		tailNumberSizeBytes := makeByteSliceAndRead(4, messageReader)

		tailNumberSizeValue := binary.BigEndian.Uint32(tailNumberSizeBytes)
		fmt.Println("Tail Number Size is:", tailNumberSizeValue)

		tailNumberValueBytes := makeByteSliceAndRead(tailNumberSizeValue, messageReader)
		tailNumberValue := string(tailNumberValueBytes)

		fmt.Println("Tail Number Value is: ", tailNumberValue)

		return &FlightMessage{}
	}

	return nil

}

func makeByteSliceAndRead(sliceSize uint32, reader *bytes.Reader) []byte {
	bytesToRead := make([]byte, sliceSize)
	reader.Read(bytesToRead)
	return bytesToRead
}
