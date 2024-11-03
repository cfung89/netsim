package netsim

import (
	"errors"
	"fmt"
	"math"

	"github.com/google/uuid"
)

type Graph struct {
	Adjacent     map[uuid.UUID][]*Edge
	Destinations []uuid.UUID
	Network      map[uuid.UUID]*Node
}

type Edge struct {
	Node   *Node // Destination
	Weight float64
}

type CalcWeight func(*Node, *Node) float64

type Requirements struct {
	MaxDist float64
}

type Path struct {
	Nodes  []*Node
	Length float64
}

func NewGraph() *Graph {
	return &Graph{}
}

func (g *Graph) AddDestination(dest *Node) {
	g.Destinations = append(g.Destinations, dest.ID)
}

func (g *Graph) AddNode(from *Node, to *Node, weight float64) error {
	if weight <= 0 {
		return errors.New("Inputted weight is smaller or equal to 0")
	}
	newEdge := &Edge{
		Node:   to,
		Weight: weight,
	}
	_, ok := g.Adjacent[from.ID]
	if ok {
		g.Adjacent[from.ID] = append(g.Adjacent[from.ID], newEdge)
	} else {
		g.Adjacent[from.ID] = []*Edge{newEdge}
	}

	return nil
}

func NetToGraph(network *Network, req Requirements, calcWeight CalcWeight) (*Graph, error) {
	if len(network.Nodes) == 0 {
		return nil, errors.New("No nodes in network.")
	}
	if len(network.Sinks) == 0 {
		return nil, errors.New("No sink nodes in network.")
	}

	graph := &Graph{
		Adjacent:     make(map[uuid.UUID][]*Edge, 0),
		Destinations: make([]uuid.UUID, len(network.Sinks)),
		Network:      make(map[uuid.UUID]*Node, 0),
	}
	for i, n := range network.Sinks {
		graph.Destinations[i] = n.ID
	}

	arrNodes := append(network.Nodes, network.Sinks...)
	for i, lnode := range arrNodes {
		for j, rnode := range arrNodes {
			if j > i && Dist(lnode.Properties.Location, rnode.Properties.Location) <= req.MaxDist {
				// Edge for left node
				ledge := &Edge{
					Node:   rnode,
					Weight: calcWeight(lnode, rnode),
				}

				// Edge for right node
				redge := &Edge{
					Node:   lnode,
					Weight: calcWeight(rnode, lnode),
				}

				lnode.Neighbours = append(lnode.Neighbours, rnode)
				graph.Adjacent[lnode.ID] = append(graph.Adjacent[lnode.ID], ledge)
				graph.Adjacent[rnode.ID] = append(graph.Adjacent[rnode.ID], redge)
			}
		}
		graph.Network[lnode.ID] = lnode
	}

	return graph, nil
}

func (g *Graph) Stdout() string {
	var s string
	for i, n := range g.Adjacent {
		var temp string
		for _, e := range n {
			temp += fmt.Sprintf(" ( %s, %.2f ) ", e.Node.ID, e.Weight)
		}
		s += fmt.Sprintf("%s: [ %s ]\n", i, temp)
	}
	fmt.Printf(s)
	return s
}

// Calculates the Euclidean distance between two points
func Dist(loc1 [3]float64, loc2 [3]float64) float64 {
	sum := math.Pow(loc1[0]-loc2[0], 2) + math.Pow(loc1[1]-loc2[1], 2) + math.Pow(loc1[2]-loc2[2], 2)
	return math.Sqrt(sum)
}
