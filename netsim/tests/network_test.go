package netsim_test

import (
	"math/rand/v2"
	"testing"

	"github.com/cfung89/sharp/netsim"
)

func TestNetwork(t *testing.T) {
	length := 100

	nodes := make([]*netsim.Node, length)
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
		nodes[i] = node
	}

	n := netsim.NewNetwork(nodes)
	tot := len(n.Sinks) + len(n.Nodes)
	if tot != length {
		t.Errorf("Incorrect network size.")
	}
	n.Stdout()

	node, err := netsim.NewNode("Sensor", rand.IntN(101), [3]float64{rand.Float64() * 1000, rand.Float64() * 1000, rand.Float64() * 1000})
	if err != nil {
		t.Errorf(err.Error())
	}
	n.AddNode(node)
	tot = len(n.Sinks) + len(n.Nodes)
	if tot != length+1 {
		t.Errorf("Incorrect network size.")
	}
	found := false
	for _, item := range n.Nodes {
		if item.ID == node.ID {
			found = true
			break
		}
	}
	if found == false {
		t.Errorf("Added node not found.")
	}

	tempNode := n.Nodes[0]
	n.RemoveNode(tempNode)
	tot = len(n.Sinks) + len(n.Nodes)
	if tot != length {
		t.Errorf("Incorrect network size.")
	}
	found = false
	for _, item := range n.Nodes {
		if item.ID == tempNode.ID {
			found = true
			break
		}
	}
	if found == true {
		t.Errorf("Node not removed.")
	}

	repetitions := 10

	// Test adding and removing nodes
	for rep := 0; rep < repetitions; rep++ {

		node, err = netsim.NewNode("Sensor", rand.IntN(101), [3]float64{rand.Float64() * 1000, rand.Float64() * 1000, rand.Float64() * 1000})
		if err != nil {
			t.Errorf(err.Error())
		}
		n.AddNode(node)
		tot = len(n.Sinks) + len(n.Nodes)
		if tot != length+1 {
			t.Errorf("Incorrect network size.")
		}
		found := false
		for _, item := range n.Nodes {
			if item.ID == node.ID {
				found = true
				break
			}
		}
		if found == false {
			t.Errorf("Added node not found.")
		}

		n.RemoveNode(node)
		tot = len(n.Sinks) + len(n.Nodes)
		if tot != length {
			t.Errorf("Incorrect network size.")
		}
		found = false
		for _, item := range n.Nodes {
			if item.ID == node.ID {
				found = true
				break
			}
		}
		if found == true {
			t.Errorf("Node not removed.")
		}

		var sink *netsim.Node
		sink, err = netsim.NewNode("Sink", rand.IntN(101), [3]float64{rand.Float64() * 1000, rand.Float64() * 1000, rand.Float64() * 1000})
		if err != nil {
			t.Errorf(err.Error())
		}
		n.AddNode(sink)
		tot = len(n.Sinks) + len(n.Nodes)
		if tot != length+1 {
			t.Errorf("Incorrect network size.")
		}
		found = false
		for _, item := range n.Sinks {
			if item.ID == sink.ID {
				found = true
				break
			}
		}
		if found == false {
			t.Errorf("Added sink node not found.")
		}

		n.RemoveNode(sink)
		tot = len(n.Sinks) + len(n.Nodes)
		if tot != length {
			t.Errorf("Incorrect network size.")
		}
		found = false
		for _, item := range n.Sinks {
			if item.ID == sink.ID {
				found = true
				break
			}
		}
		if found == true {
			t.Errorf("Sink node not removed.")
		}
	}
}
