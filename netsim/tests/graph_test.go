package netsim_test

import (
	"fmt"
	"math/rand/v2"
	"testing"

	"github.com/cfung89/sharp/netsim"
)

func TestGraph(t *testing.T) {
	length := 100

	nodes := make([]*netsim.Node, length)
	for i := range nodes {
		random := rand.IntN(3)
		var status netsim.NodeType
		if random < 2 {
			status = netsim.Sensor
		} else {
			status = netsim.Sink
		}

		node, err := netsim.NewNode(status, rand.IntN(101), [3]float64{rand.Float64() * 1000, rand.Float64() * 1000, rand.Float64() * 1000})
		if err != nil {
			t.Errorf(err.Error())
		}
		nodes[i] = node
	}

	network := netsim.NewNetwork(nodes)
	tot := len(network.Sinks) + len(network.Nodes)
	if tot != length {
		t.Errorf("Incorrect network size.")
	}

	req := netsim.Requirements{
		MaxDist: 100,
	}
	g, err := netsim.NetToGraph(network, req)
	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Println(len(g.Adjacent))
	g.Stdout()
}
