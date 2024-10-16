package netsim

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
	validDataTypes := []string{"ClientHello", "ClientCom", "ClientClosed"}
	valid := false
	for _, n := range validDataTypes {
		if n == dataType {
			valid = true
		}
	}

	if !valid {
		return nil, errors.New("Invalid data type.")
	} else if len(data) == 0 {
		return nil, errors.New("Invalid data length.")
	}

	newPacket := &Packet{
		Data:   data,
		Length: len(data),
		Type:   dataType,
	}

	return newPacket, nil
}
