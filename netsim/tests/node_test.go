package netsim_test

import (
	"math/rand/v2"
	"testing"
	"time"

	"github.com/cfung89/sharp/netsim"
)

func TestNode(t *testing.T) {
	length := 10
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
		node.Stdout()
		nodes[i] = node
	}
}

func TestSendRead(t *testing.T) {
	node1, _ := netsim.NewNode(netsim.Sensor, rand.IntN(101), [3]float64{rand.Float64() * 1000, rand.Float64() * 1000, rand.Float64() * 1000})
	node2, _ := netsim.NewNode(netsim.Sensor, rand.IntN(101), [3]float64{rand.Float64() * 1000, rand.Float64() * 1000, rand.Float64() * 1000})
	node1.Neighbours = append(node1.Neighbours, node2)
	node2.Neighbours = append(node2.Neighbours, node1)

	go func() {
		packet, err := netsim.NewPacket([]byte("test"), netsim.NetworkCom)
		if err != nil {
			t.Errorf(err.Error())
		}

		packet.Headers.Path.Nodes = append(packet.Headers.Path.Nodes, node2)
		sendErr := node1.Send(packet)
		if sendErr != nil {
			t.Errorf(sendErr.Error())
		}
		time.Sleep(2 * time.Second)

	}()

	// Close
	go func() {

		time.Sleep(2 * time.Second)
		packet, err := netsim.NewPacket([]byte("Close"), netsim.NetworkClose)
		if err != nil {
			t.Errorf(err.Error())
		}
		packet.Headers.Path.Nodes = append(packet.Headers.Path.Nodes, node2)
		sendErr := node1.Send(packet)
		if sendErr != nil {
			t.Errorf(sendErr.Error())
		}
	}()

	node2.Read()

}
