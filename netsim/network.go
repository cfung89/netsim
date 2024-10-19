package netsim

import (
	"errors"
	"fmt"
	"strings"
)

type Network struct {
	Sinks []*Node
	Nodes []*Node
}

func NewNetwork(nodes []*Node) *Network {
	nw := &Network{}
	for _, n := range nodes {
		nw.AddNode(n)
	}
	return nw
}

func (network *Network) AddSink(sink *Node) error {
	if sink.Properties.Type != "Sink" {
		return errors.New("Not a sink node")
	}
	network.Sinks = append(network.Sinks, sink)
	return nil
}

func (network *Network) AddNode(node *Node) {
	if node.Properties.Type == "Sink" {
		network.AddSink(node)
	} else {
		network.Nodes = append(network.Nodes, node)
	}
}

func (network *Network) RemoveNode(node *Node) error {
	if node.Properties.Type == "Sink" {
		for i, n := range network.Sinks {
			if n == node {
				network.Sinks = append(network.Sinks[:i], network.Sinks[i+1:]...)
				return nil
			}
		}
		return errors.New("Cannot find node to remove.")
	} else if node.Properties.Type == "Sensor" {
		for i, n := range network.Nodes {
			if n == node {
				network.Nodes = append(network.Nodes[:i], network.Nodes[i+1:]...)
				return nil
			}
		}
		return errors.New("Cannot find node to remove.")
	} else {
		return errors.New("Invalid node type.")
	}
}

func (network *Network) Stdout() string {
	s := "Network:"
	s += fmt.Sprintf("\nSink nodes:\n")
	for i, n := range network.Sinks {
		s += fmt.Sprintf("Sink #%d:\n", i+1)
		temp := strings.Replace(n.Stdout(), "\n", "\n\t- ", 3)
		s += strings.Replace(temp, "ID:", "\t- ID:", 1)
	}

	s += fmt.Sprintf("\nSensor nodes:\n")
	for i, n := range network.Nodes {
		s += fmt.Sprintf("Node #%d:\n", i+1)
		temp := strings.Replace(n.Stdout(), "\n", "\n\t- ", 3)
		s += strings.Replace(temp, "ID:", "\t- ID:", 1)
	}
	fmt.Printf(s)
	return s
}
