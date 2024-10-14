package main

import (
	"errors"
)

type Packet struct {
	Data   []byte
	Length int
	Type   string
}

func NewEmptyPacket() *Packet {
	return &Packet{
		Data:   make([]byte, 0),
		Length: 0,
		Type:   "",
	}
}

func NewPacket(data []byte, dataType string) (*Packet, error) {
	validDataTypes := map[string]bool{
		"ClientHello":  true,
		"ClientCom":    true,
		"ClientClosed": true,
	}
	if !(validDataTypes[dataType]) {
		emptyPacket := NewEmptyPacket()
		return emptyPacket, errors.New("Invalid data type.")
	} else if len(data) == 0 {
		emptyPacket := NewEmptyPacket()
		return emptyPacket, errors.New("Invalid data length.")
	}

	newPacket := &Packet{
		Data:   data,
		Length: len(data),
		Type:   dataType,
	}

	return newPacket, nil

}
