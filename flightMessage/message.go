package flightMessage

import (
	"bytes"
	"encoding/binary"
	"errors"
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

var byteArraySizeError = errors.New("Byte array does not meet minimum size requirement")
var unexpectedHeaderError = errors.New("Got unexpected message, header bytes unrecognized")

// CreateMessageFromBinary returns a FlightMessage after decoding the information from a byte array
func CreateMessageFromBinary(messageBytes []byte) (*FlightMessage, error) {

	messageReader := bytes.NewReader(messageBytes)
	headerBytes, err := makeByteSliceAndRead(3, messageReader)
	if err != nil {
		return nil, err
	}
	fmt.Println("Header bytes are: ", headerBytes)

	if bytes.Compare(headerBytes, HEADER) == 0 {
		//TODO: Gotta be a better way to make this look cleaner
		tailNumberSizeBytes, err := makeByteSliceAndRead(4, messageReader)
		if err != nil {
			return nil, err
		}

		tailNumberSizeValue := binary.BigEndian.Uint32(tailNumberSizeBytes)
		// fmt.Println("Tail Number Size is:", tailNumberSizeValue)

		tailNumberValueBytes, err := makeByteSliceAndRead(tailNumberSizeValue, messageReader)
		if err != nil {
			return nil, err
		}
		tailNumberValue := string(tailNumberValueBytes)
		// fmt.Println("Tail Number Value is: ", tailNumberValue)

		engineCountBytes, err := makeByteSliceAndRead(4, messageReader)
		if err != nil {
			return nil, err
		}
		engineCount := binary.BigEndian.Uint32(engineCountBytes)
		// fmt.Println("Engine Count is: ", engineCount)

		engineNameSizeBytes, err := makeByteSliceAndRead(4, messageReader)
		if err != nil {
			return nil, err
		}
		engineNameSize := binary.BigEndian.Uint32(engineNameSizeBytes)
		// fmt.Println("Engine Name Size is: ", engineNameSize)

		engineNameValueBytes, err := makeByteSliceAndRead(engineNameSize, messageReader)
		if err != nil {
			return nil, err
		}
		engineNameValue := string(engineNameValueBytes)
		// fmt.Println("Engine Name is: ", engineNameValue)

		latitudeBytes, err := makeByteSliceAndRead(8, messageReader)
		if err != nil {
			return nil, err
		}
		latitude, err := Float64frombytes(latitudeBytes)
		if err != nil {
			return nil, err
		}
		// fmt.Println("Latitude is: ", latitude)

		longitudeBytes, err := makeByteSliceAndRead(8, messageReader)
		if err != nil {
			return nil, err
		}
		longitude, err := Float64frombytes(longitudeBytes)
		if err != nil {
			return nil, err
		}
		// fmt.Println("Longitude is: ", longitude)

		altitudeBytes, err := makeByteSliceAndRead(8, messageReader)
		if err != nil {
			return nil, err
		}
		altitude, err := Float64frombytes(altitudeBytes)
		if err != nil {
			return nil, err
		}
		// fmt.Println("Altitude is: ", altitude)

		temperatureBytes, err := makeByteSliceAndRead(8, messageReader)
		if err != nil {
			return nil, err
		}
		temperature, err := Float64frombytes(temperatureBytes)
		if err != nil {
			return nil, err
		}
		// fmt.Println("Temperature is: ", temperature)

		return &FlightMessage{
			tail_number:  tailNumberValue,
			engine_count: int(engineCount),
			engine_name:  engineNameValue,
			latitude:     *latitude,
			longitude:    *longitude,
			altitude:     *altitude,
			temperature:  *temperature}, nil
	}

	return nil, unexpectedHeaderError

}

// makeByteSliceAndRead takes in a uint32 and byte Reader, creates a byte array of size sliceSize and reads from the Reader into the byte array
func makeByteSliceAndRead(sliceSize uint32, reader *bytes.Reader) ([]byte, error) {
	bytesToRead := make([]byte, sliceSize)
	_, err := reader.Read(bytesToRead)
	if err != nil {
		fmt.Printf("Got error from reader.Read, error is: %s", err)
		return nil, err
	}

	return bytesToRead, nil
}

// Float64frombytes takes in a byte array and transforms it into a float64. This function assumes the byte array is BigEndian
func Float64frombytes(bytes []byte) (*float64, error) {
	if len(bytes) < 7 {
		return nil, byteArraySizeError
	}
	bits := binary.BigEndian.Uint64(bytes)
	float := math.Float64frombits(bits)
	return &float, nil
}
