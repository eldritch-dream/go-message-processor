package flightMessage

import (
	"bytes"
	"encoding/binary"
	"io"
	"math"
	"testing"
)

func TestMakeByteSliceAndRead(t *testing.T) {

	var testString = "CORGI"

	var testBytes []byte
	testBytes = append(testBytes, testString...)
	byteReader := bytes.NewReader(testBytes)

	c, err := makeByteSliceAndRead(5, byteReader)
	if err != nil {
		t.Errorf("Got error %s but expected no error", err)
	}
	corgi := string(c)
	if corgi != "CORGI" {
		t.Errorf("Got %s, but expected CORGI", c)
	}

	_, err = makeByteSliceAndRead(1, byteReader)
	if err != nil {
		if err != io.EOF {
			t.Errorf("Got error %s but expected io.EOF error", err)
		}
	}
}

func TestFloat64frombytes(t *testing.T) {
	testFloat := 1234.567890
	bits := math.Float64bits(testFloat)
	bytes := make([]byte, 8)
	binary.BigEndian.PutUint64(bytes, bits)

	convertedFloat, err := Float64frombytes(bytes)
	if err != nil {
		t.Errorf("Expected successful conversion but got error: %s", err)
	}
	if *convertedFloat != testFloat {
		t.Errorf("Got %f but expected %f", *convertedFloat, testFloat)
	}

	var emptyBytes []byte
	expectedError := byteArraySizeError
	_, err = Float64frombytes(emptyBytes)
	if err != nil {
		if err != expectedError {
			t.Errorf("Expected error %s, to equal %s", err, expectedError)
		}
	}
}
