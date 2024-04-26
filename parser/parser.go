package parser

import (
	"encoding/hex"
	"fmt"
	"strconv"

	consolidated076can "github.com/Ion-Mobility/can_parser_example/candbc"
	"go.einride.tech/can"
)

type Message struct {
	BikeID    string
	Timestamp string
	CanID     string
	Length    float64
	Data      string
}

func Parse(msg Message) (map[string]interface{}, error) {
	id, err := hexStringToUint32(msg.CanID)
	if err != nil {
		return nil, fmt.Errorf("invalid can ID: %v", err)
	}

	data, err := hexStringToBytes(msg.Data)
	if err != nil {
		return nil, fmt.Errorf("invalid data: %v", err)
	}

	// Encode can data
	f := can.Frame{ID: id, Length: uint8(msg.Length), Data: data, IsExtended: true}
	md := consolidated076can.Messages()
	ms, err := md.UnmarshalFrame(f)
	if err != nil {
		return nil, fmt.Errorf("unmarshal frame: %v", err)
	}

	// Map to key value
	mapping := make(map[string]interface{})
	signals := ms.Descriptor().Signals
	for _, s := range signals {
		d := ms.Frame().Data
		valueDescription, hasValueDescription := s.UnmarshalValueDescription(d)
		switch {
		case hasValueDescription:
			mapping[s.Name] = valueDescription
		case s.Length == 1: // bool
			val := s.UnmarshalBool(d)
			mapping[s.Name] = val
		case s.IsSigned: // signed
			val := s.UnmarshalPhysical(d)
			mapping[s.Name] = val
		default: // unsigned
			val := s.UnmarshalPhysical(d)
			mapping[s.Name] = val
		}
	}

	return mapping, nil
}

func hexStringToUint32(hexString string) (uint32, error) {
	// Convert hexadecimal string to integer
	result, err := strconv.ParseUint(hexString, 16, 32)
	if err != nil {
		return 0, err
	}
	return uint32(result), nil
}

func hexStringToBytes(hexString string) ([8]byte, error) {
	var result [8]byte
	// DecodeString returns a []byte
	bytes, err := hex.DecodeString(hexString)
	if err != nil {
		return result, err
	}
	// Copy bytes into result array
	copy(result[:], bytes)
	return result, nil
}
