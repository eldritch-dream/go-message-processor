package flightMessage

import (
	"bytes"
	"encoding/binary"
	"math"
	"testing"
)

func TestMakeByteSliceAndRead(t *testing.T) {

	var testString = "CORGI"

	var testBytes []byte
	testBytes = append(testBytes, testString...)
	byteReader := bytes.NewReader(testBytes)

	c := string(makeByteSliceAndRead(5, byteReader))
	if c != "CORGI" {
		t.Errorf("Got %s, but expected CORGI", c)
	}
}

func TestFloat64frombytes(t *testing.T) {
	testFloat := 1234.567890
	bits := math.Float64bits(testFloat)
	bytes := make([]byte, 8)
	binary.BigEndian.PutUint64(bytes, bits)

	convertedFloat := Float64frombytes(bytes)
	if convertedFloat != testFloat {
		t.Errorf("Got %f but expected %f", convertedFloat, testFloat)
	}
}

