package netsim

import (
	"errors"
)

type Network struct {
	Sink  *Node // not in Nodes array
	Nodes []*Node
}

func NewNetwork(nodes []*Node) *Network {
	return &Network{
		Nodes: nodes,
	}
}

func (network *Network) SetSink(sink *Node) {
	network.Sink = sink
}

func (network *Network) AddNode(node *Node) {
	network.Nodes = append(network.Nodes, node)
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
