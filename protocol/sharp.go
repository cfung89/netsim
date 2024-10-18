package protocol

import (
	"github.com/cfung89/sharp/netsim"
	// "io"
	"fmt"
	"math"
	"net"
)

type Router struct {
	from net.Conn
	to   net.Conn
}

func main() {
	network := netsim.NewNetwork(make([]*netsim.Node, 0))
	fmt.Println(network)
}

// Calculates the Euclidean distance between two points
func Dist(loc1 [3]float64, loc2 [3]float64) float64 {
	sum := math.Pow(loc1[0]-loc2[0], 2) + math.Pow(loc1[1]-loc2[1], 2) + math.Pow(loc1[2]-loc2[2], 2)
	return math.Sqrt(sum)
}

// go func() {
// 	io.Copy(p.from, p.to)
// 	cancel()
// }
