package flightMessage

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
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

		engineCountBytes := makeByteSliceAndRead(4, messageReader)
		engineCount := binary.BigEndian.Uint32(engineCountBytes)
		fmt.Println("Engine Count is: ", engineCount)

		engineNameSizeBytes := makeByteSliceAndRead(4, messageReader)
		engineNameSize := binary.BigEndian.Uint32(engineNameSizeBytes)
		fmt.Println("Engine Name Size is: ", engineNameSize)

		engineNameValueBytes := makeByteSliceAndRead(engineNameSize, messageReader)
		engineNameValue := string(engineNameValueBytes)
		fmt.Println("Engine Name is: ", engineNameValue)

		latitudeBytes := makeByteSliceAndRead(8, messageReader)
		latitude := Float64frombytes(latitudeBytes)
		fmt.Println("Latitude is: ", latitude)

		longitudeBytes := makeByteSliceAndRead(8, messageReader)
		longitude := Float64frombytes(longitudeBytes)
		fmt.Println("Longitude is: ", longitude)

		altitudeBytes := makeByteSliceAndRead(8, messageReader)
		altitude := Float64frombytes(altitudeBytes)
		fmt.Println("Altitude is: ", altitude)

		temperatureBytes := makeByteSliceAndRead(8, messageReader)
		temperature := Float64frombytes(temperatureBytes)
		fmt.Println("Temperature is: ", temperature)

		return &FlightMessage{tail_number: tailNumberValue, engine_count: int(engineCount), engine_name: engineNameValue, latitude: latitude, longitude: longitude, altitude: altitude, temperature: temperature}
	}

	return nil

}

func makeByteSliceAndRead(sliceSize uint32, reader *bytes.Reader) []byte {
	bytesToRead := make([]byte, sliceSize)
	reader.Read(bytesToRead)
	return bytesToRead
}

func Float64frombytes(bytes []byte) float64 {
	bits := binary.BigEndian.Uint64(bytes)
	float := math.Float64frombits(bits)
	return float
}
