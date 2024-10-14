package main

import (
	"errors"
)

type Network struct {
	Nodes []*Node
}

func NewNetwork() *Network {
	return nil
}

func AddNode(node *Node) error {
	return nil
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
