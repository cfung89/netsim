package netsim

import (
	"errors"
	"github.com/google/uuid"
)

type Node struct {
	Type     string
	ID       uuid.UUID
	Battery  int
	Location [3]float64
	Children []*Node
}

func NewNode(status string, initBattery int, location [3]float64) (*Node, error) {
	if status != "Sensor" && status != "Sink" {
		emptyNode := &Node{
			Type:     status,
			Battery:  initBattery,
			Location: location,
		}
		return emptyNode, errors.New("Unknown status")
	}

	newNode := &Node{
		ID:       uuid.New(),
		Type:     status,
		Battery:  initBattery,
		Location: location,
	}

	return newNode, nil
}
