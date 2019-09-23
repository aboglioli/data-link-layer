package implementation

import (
	"github.com/aboglioli/data-link-layer/frame"
	"github.com/aboglioli/data-link-layer/packet"
	"github.com/aboglioli/data-link-layer/physical"
	"github.com/aboglioli/data-link-layer/protocol"
)

type utopianSimplex struct {
	done chan bool
}

func UtopianSimplex() *utopianSimplex {
	return &utopianSimplex{
		done: make(chan bool),
	}
}

func (u *utopianSimplex) StartReceiver() {
	phy := physical.TCPServer()

	prot := protocol.NewGeneric(phy)

	for {
		var f frame.Frame
		prot.FromPhysicalLayer(&f)
		prot.ToNetworkLayer(&f.Info)
	}
}

func (u *utopianSimplex) StartSender() {
	phy := physical.TCPClient()

	prot := protocol.NewGeneric(phy)

	for {
		var f frame.Frame
		var pk packet.Packet
		prot.FromNetworkLayer(&pk)
		f.Info = pk
		prot.ToPhysicalLayer(&f)
	}
}

func (u *utopianSimplex) Stop() {
	u.done <- true
}

func (u *utopianSimplex) Wait() bool {
	return <-u.done
}
