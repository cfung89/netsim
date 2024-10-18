package netsim

import (
	"errors"
	"github.com/google/uuid"
)

type Graph struct {
	Adjacent     map[uuid.UUID][]*Edge
	Destinations []uuid.UUID
}

type Edge struct {
	Node   *Node // Destination
	Weight float64
}

// type CalcWeight func(*Node, *Node) float64

type Requirements struct {
	DistThreshold float64
}

func NewGraph() *Graph {
	return &Graph{}
}

func (g *Graph) AddDestination(dest *Node) {
    g.Destinations = append(g.Destinations, dest.ID)
}

func (g *Graph) AddNode(from *Node, to *Node, weight float64) {
	newEdge := &Edge{
		Node:   to,
		Weight: weight,
	}
    val, ok := g.Adjacent[from.ID]
    if ok {
        g.Adjacent[from.ID] = append(g.Adjacent[from.ID], newEdge)
    } else {
        g.Adjacent[from.ID] = []*Edge{newEdge}
    }
}

func NetToGraph(network *Network, req *Requirements) (*Graph, error) {
	if len(network.Nodes) == 0 {
		return nil, errors.New("No nodes in network.")
	}
	if len(network.Sinks) == 0 {
		return nil, errors.New("No sink nodes in network.")
	}

	graph := &Graph{
		Adjacent:     make(map[uuid.UUID][]*Edge),
		Destinations: network.Sinks,
	}

	arrNodes = append(network.Nodes, network.Sinks...)
	for i, lnode := range arrNodes {
		for j, rnode := range arrNodes {
			if j > i && Dist(lnode.Location, rnode.Location) <= req.DistThreshold {
				ledge = &Edge{
					Node:   rnode,
					Weight: 0,
					// Weight: calcWeight(lnode, rnode),
				}
				graph.Adjacent[lnode.ID] = append(graph.Adjacent[lnode.ID], edge)

				redge = &Edge{
					Node:   lnode,
					Weight: 0,
					// Weight: calcWeight(rnode, lnode),
				}
				graph.Adjacent[rnode.ID] = append(graph.Adjacent[rnode.ID], edge)
			}
		}
		lnode.Neighbours = graph.Adjacent[lnode.ID]
	}

	return graph, nil
}
