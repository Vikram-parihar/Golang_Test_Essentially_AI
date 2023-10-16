package main

import (
    "encoding/binary"
    "fmt"
)

type DecodedData struct {
    Short1        uint16
    TwelveChars   string
    SingleByte    byte
    EightChars    string
    Short2        uint16
    FifteenChars  string
    Long          uint32
}
// function to decode the packet data 
func decodePacket(packet []byte) (DecodedData, error) {
    // Error condition to check if the byte slice is lesser or greater than 44 bytes.
    if len(packet) != 44 {
        return DecodedData{}, fmt.Errorf("Invalid packet size. Expected 44 bytes, but got %d", len(packet))
    }

    decoded := DecodedData{}
    offset := 0

    
    decoded.Short1 = binary.BigEndian.Uint16(packet[offset : offset+2])
    offset += 2

    decoded.TwelveChars = string(packet[offset : offset+12])
    offset += 12


    decoded.SingleByte = packet[offset]
    offset++

   
    decoded.EightChars = string(packet[offset : offset+8])
    offset += 8


    decoded.Short2 = binary.BigEndian.Uint16(packet[offset : offset+2])
  
    decoded.FifteenChars = string(packet[offset : offset+15])
    offset += 15

    decoded.Long = binary.BigEndian.Uint32(packet[offset : offset+4])

    return decoded, nil
}

func main() {
    hexSequence := []byte{
        0x04, 0xD2, 0x6B, 0x65, 0x65, 0x70, 0x64, 0x65, 0x63, 0x6F, 0x64, 0x69,
        0x6E, 0x67, 0x38, 0x64, 0x6F, 0x6E, 0x74, 0x73, 0x74, 0x6F, 0x70, 0x03,
        0x15, 0x63, 0x6F, 0x6E, 0x67, 0x72, 0x61, 0x74, 0x75, 0x6C, 0x61, 0x74,
        0x69, 0x6F, 0x6E, 0x73, 0x07, 0x5B, 0xCD, 0x15,
    }

    decoded, err := decodePacket(hexSequence)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Printf("Decoded struct: %+v\n", decoded)
    }
}
