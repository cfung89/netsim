package netsim

import (
	"errors"
	"fmt"
)

type DataType string

type Headers struct {
	Length int
	Type   DataType
	Path   *Path
}

type Payload struct {
	Data []byte
}

type Packet struct {
	Headers *Headers
	Payload *Payload
}

const (
	NetworkInit   DataType = "NetworkInit"
	NetworkCom    DataType = "NetworkCom"
	NetworkUpdate DataType = "NetworkUpdate"
	NetworkClose  DataType = "NetworkClose"
)

var validDataTypes = map[DataType]struct{}{
	NetworkInit:   {},
	NetworkCom:    {},
	NetworkUpdate: {},
	NetworkClose:  {},
}

func IsValidDataType(dataType DataType) bool {
	_, ok := validDataTypes[dataType]
	return ok
}

func NewPacket(data []byte, dataType DataType) (*Packet, error) {
	if !IsValidDataType(dataType) {
		return nil, errors.New("Invalid data type.")
	} else if len(data) == 0 {
		return nil, errors.New("Invalid data length.")
	}

	newHeaders := &Headers{
		Length: len(data),
		Type:   dataType,
		Path: &Path{
			Nodes:  make([]*Node, 0),
			Length: 0,
		},
	}
	newPayload := &Payload{
		Data: data,
	}
	newPacket := &Packet{
		Headers: newHeaders,
		Payload: newPayload,
	}

	return newPacket, nil
}

func (packet *Packet) Stdout() string {
	var s string = "Headers:\n"
	s += fmt.Sprintf("\t- Data Length: %d\n\t- Packet Type: %s\n\t- Path: %s\n\t- Path Length: %.2f\n\tPath: [\n", packet.Headers.Length, packet.Headers.Type, packet.Headers.Path.Nodes, packet.Headers.Path.Length)
	for _, n := range packet.Headers.Path.Nodes {
		s += fmt.Sprintf("\t\t%s,\n", n.ID)
	}
	s += fmt.Sprintf("\t\t]\n\t- Path Length: %.2f\n", packet.Headers.Path.Length)
	s += fmt.Sprintf("Payload:\n\t- %s\n\n", packet.Payload.Data)
	fmt.Printf(s)
	return s
}
