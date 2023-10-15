package main

import (
	"encoding/binary"
	"fmt"
	"strings"
)

type DecodedStruct struct {
	Short1 uint16
	Chars1 string
	Byte1  uint8
	Chars2 string
	Short2 uint16
	Chars3 string
	Long1  uint32
}

func decodePacket(packet []byte) (*DecodedStruct, error) {
	if len(packet) != 44 {
		return nil, fmt.Errorf("invalid packet size, expected 44 bytes")
	}

	var decoded DecodedStruct

	// Unpack the data based on the decoding map
	decoded.Short1 = binary.BigEndian.Uint16(packet[0:2])
	decoded.Chars1 = strings.TrimRight(string(packet[2:14]), "\x00")
	decoded.Byte1 = packet[14]
	decoded.Chars2 = strings.TrimRight(string(packet[15:23]), "\x00")
	decoded.Short2 = binary.BigEndian.Uint16(packet[23:25])
	decoded.Chars3 = strings.TrimRight(string(packet[25:40]), "\x00")
	decoded.Long1 = binary.BigEndian.Uint32(packet[40:44])

	return &decoded, nil
}

func main() {
	packet := []byte{
		0x04, 0xD2, 0x6B, 0x65, 0x65, 0x70, 0x64, 0x65, 0x63, 0x6F, 0x64, 0x69, 0x6E, 0x67,
		0x38, 0x64, 0x6F, 0x6E, 0x74, 0x73, 0x74, 0x6F, 0x70, 0x03, 0x15, 0x63, 0x6F, 0x6E,
		0x67, 0x72, 0x61, 0x74, 0x75, 0x6C, 0x61, 0x74, 0x69, 0x6F, 0x6E, 0x73, 0x07, 0x5B, 0xCD, 0x15,
	}

	decodedStruct, err := decodePacket(packet)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("%+v\n", *decodedStruct)
}
