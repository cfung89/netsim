package protocol

import (
	"github.com/cfung89/sharp/netsim"
	// "io"
	"fmt"
	"net"
)

type Router struct {
	from net.Conn
	to   net.Conn
}

func main() {
	network := netsim.NewNetwork()
	fmt.Println(network)
}

// go func() {
// 	io.Copy(p.from, p.to)
// 	cancel()
// }
