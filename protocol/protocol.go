package protocol

import (
	"fmt"
	"math/rand/v2"

	"github.com/cfung89/sharp/netsim"
)

func calcWeight(node1 *netsim.Node, node2 *netsim.Node) float64 {
	return netsim.Dist(node1.Properties.Location, node2.Properties.Location)
}

func Protocol(length int) error {
	nodes := make([]*netsim.Node, length)
	for i := range nodes {
		random := rand.IntN(10)
		// random := rand.IntN(20)

		var status netsim.NodeType
		if random < 1 {
			status = netsim.Sink
		} else {
			status = netsim.Sensor
		}

		node, err := netsim.NewNode(status, rand.IntN(101), [3]float64{rand.Float64() * 1000, rand.Float64() * 1000, rand.Float64() * 1000})
		if err != nil {
			return err
		}
		nodes[i] = node
	}

	network := netsim.NewNetwork(nodes)
	req := netsim.Requirements{
		MaxDist: 100,
	}
	g, convErr := netsim.NetToGraph(network, req, calcWeight)
	if convErr != nil {
		return convErr
	}

	// Get weights
	// for id, arrEdge := range g.Adjacent {
	// 	for _, edge := range arrEdge {
	// 		edge.Weight = netsim.Dist(g.Network[id].Properties.Location, edge.Node.Properties.Location)
	// 	}
	// }

	// for _, arrEdge := range g.Adjacent {
	// 	for _, edge := range arrEdge {
	// 		fmt.Println(edge.Weight)
	// 	}
	// }

	// Calculate paths
	source := network.Nodes[rand.IntN(len(network.Nodes))]
	paths, err := Dijkstra(g, source.ID)
	if err != nil {
		return err
	}

	for i, n := range paths {
		fmt.Println(i, n.Nodes, n.Length)
	}
	return nil
}
