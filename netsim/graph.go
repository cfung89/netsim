package netsim

import (
	"errors"
)

type Graph struct {
	Start       *Node
	Destination *Node
}

func (network *Network) NetToGraph(req *Requirements) (*Graph, error) {
	if len(network.Nodes) == 0 {
		return nil, errors.New("No nodes in network.")
	}

	graph := &Graph{
		Destination: network.Sink,
	}

	for _, lnode := range network.Nodes {
		for _, rnode := range network.Nodes {
			if lnode == rnode || Dist(lnode.Location, rnode.Location) > req.DistThreshold {
				continue
			}
			lnode.Children = append(lnode.Children, rnode)
		}
	}

	return graph, nil
}
