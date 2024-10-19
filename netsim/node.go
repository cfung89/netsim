package netsim

import (
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
)

type NodeProperties struct {
	Type     string
	Battery  int
	Location [3]float64
}

type Node struct {
	ID         uuid.UUID
	Channel    chan *Packet
	Properties *NodeProperties
	Neighbours []*Node
}

func NewNode(status string, initBattery int, location [3]float64) (*Node, error) {
	if initBattery > 100 || initBattery < 0 {
		return nil, errors.New("Invablid initial battery value.")
	}
	if status != "Sensor" && status != "Sink" {
		return nil, errors.New("Unknown status.")
	}

	newNodeProperties := &NodeProperties{
		Type:     status,
		Battery:  initBattery,
		Location: location,
	}
	newNode := &Node{
		ID:         uuid.New(),
		Channel:    make(chan *Packet),
		Properties: newNodeProperties,
		Neighbours: make([]*Node, 0),
	}
	// go newNode.Read()
	return newNode, nil
}

func (node *Node) Read() {
	for {
		if len(node.Channel) != 0 {
			msg := <-node.Channel
			log.Println(msg)
		}
	}
}

func (node *Node) Send(packet *Packet, dest uuid.UUID) error {
	for _, n := range node.Neighbours {
		if n.ID == dest {
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
