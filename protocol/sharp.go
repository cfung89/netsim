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

// go func() {
// 	io.Copy(p.from, p.to)
// 	cancel()
// }
