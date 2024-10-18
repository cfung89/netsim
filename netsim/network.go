package netsim

import (
	"errors"
)

type Network struct {
	Sinks  []*Node
	Nodes []*Node
}

func NewNetwork() *Network {
	return &Network{}
}

func (network *Network) SetSink(sink *Node) error {
    if sink.Properties.Type != "Sink" {
        return errors.New("Not a sink node")
    }
	network.Sinks = append(network.Sinks, sink)
    return nil
}

func (network *Network) AddNode(node *Node) {
    if node.Properties.Type == "Sink" {
        network.SetSink(node)
    } else {
        network.Nodes = append(network.Nodes, node)
    }
}

func (network *Network) RemoveNode(node *Node) error {
	for i, n := range network.Nodes {
		if n == node {
			network.Nodes = append(network.Nodes[:i], network.Nodes[i+1:]...)
			return nil
		}
	}
	return errors.New("Cannot find node to remove")
}
