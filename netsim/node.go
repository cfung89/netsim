package netsim

import (
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
)

type NodeType string

const (
	Sink   NodeType = "Sink"
	Sensor NodeType = "Sensor"
)

var validNodeTypes = map[NodeType]struct{}{
	Sink:   {},
	Sensor: {},
}

func isValidNodeType(nodeType NodeType) bool {
	_, ok := validNodeTypes[nodeType]
	return ok
}

type NodeProperties struct {
	Type     NodeType
	Battery  int
	Location [3]float64
}

type Node struct {
	ID         uuid.UUID
	Channel    chan *Packet
	Properties *NodeProperties
	Neighbours []*Node
}

func NewNode(status NodeType, initBattery int, location [3]float64) (*Node, error) {
	if initBattery > 100 || initBattery < 0 {
		return nil, errors.New("Invablid initial battery value.")
	}
	if !isValidNodeType(status) {
		return nil, errors.New("Unknown node type.")
	}

	newNodeProperties := &NodeProperties{
		Type:     status,
		Battery:  initBattery,
		Location: location,
	}
	newNode := &Node{
		ID:         uuid.New(),
		Channel:    make(chan *Packet, 1),
		Properties: newNodeProperties,
		Neighbours: make([]*Node, 0),
	}
	// go newNode.Read()
	return newNode, nil
}

func (node *Node) Read() {
	for {
		select {
		case msg := <-node.Channel:
			log.Println(fmt.Sprintf("%s", msg.Payload.Data))
			if msg.Headers.Type == NetworkClose {
				return
			}
		default:
		}
	}
}

func (node *Node) Send(packet *Packet) error {
	for _, n := range node.Neighbours {
		if n.ID == packet.Headers.Path.Nodes[0].ID {
			n.Channel <- packet
			return nil
		}
	}
	return errors.New("Destination not found")
}

func (node *Node) Stdout() string {
	s := fmt.Sprintf("ID: %s\nType: %s\nBattery: %d \nLocation: x = %f; y = %f; z = %f\n\n", node.ID, node.Properties.Type, node.Properties.Battery, node.Properties.Location[0], node.Properties.Location[1], node.Properties.Location[2])
	fmt.Printf(s)
	return s
}
