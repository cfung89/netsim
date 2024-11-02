package netsim_test

import (
	"testing"

	"github.com/cfung89/sharp/netsim"
)

func TestPacket(t *testing.T) {
	packet, err := netsim.NewPacket([]byte("test"), netsim.NetworkCom)
	if err != nil {
		t.Errorf(err.Error())
	}
	packet.Stdout()
}
