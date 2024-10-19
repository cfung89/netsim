package netsim_test

import (
	"math/rand/v2"
	"testing"

	"github.com/cfung89/sharp/netsim"
)

func TestNode(t *testing.T) {
	nodes := make([]*netsim.Node, 10)
	for i := range nodes {
		random := rand.IntN(3)
		var status string
		if random < 2 {
			status = "Sensor"
		} else {
			status = "Sink"
		}

		node, err := netsim.NewNode(status, rand.IntN(101), [3]float64{rand.Float64() * 1000, rand.Float64() * 1000, rand.Float64() * 1000})
		if err != nil {
			t.Errorf(err.Error())
		}
		node.Stdout()
		nodes[i] = node
	}
}
