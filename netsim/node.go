package netsim

import (
	"errors"
	"github.com/google/uuid"
	"log"
)

type NodeProperties struct {
	Type     string
	Battery  int
	Location [3]float64
}

type Node struct {
	ID         uuid.UUID
	Channel    chan string
	Properties *NodeProperties
	Neighbours []*Node
}

func NewNode(status string, initBattery int, location [3]float64) (*Node, error) {
	if initBattery > 100 || initBattery < 0 {
	}
	if status == "Sensor" || status == "Sink" {
		newNodeProperties := &NodeProperties{
			Type:     status,
			Battery:  initBattery,
			Location: location,
		}
		newNode := &Node{
			ID:         uuid.New(),
			Channel:    make(chan string),
			Properties: newNodeProperties,
			Neighbours: []*Node,
		}
		go newNode.Read()
		return newNode, nil
	}

	return nil, errors.New("Unknown status")

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
	for i, n := range node.Neighbours {
		if i == dest {
			n.Channel <- packet
			return nil
		}
	}
	return errors.New("Destination not found")
}
