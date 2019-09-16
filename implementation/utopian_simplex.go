package implementation

import (
	"fmt"

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
	fmt.Println("Receptor iniciado")

	phy := physical.TCPServer()

	prot := protocol.NewGeneric(phy)

	for {
		var f frame.Frame

		prot.FromPhysicalLayer(&f)
		fmt.Println("FromPhysicalLayer", f)

		prot.ToNetworkLayer(&f.Info)
		fmt.Println("ToNetworkLayer", f.Info)
	}
}

func (u *utopianSimplex) StartSender() {
	fmt.Println("Emisor iniciado")

	phy := physical.TCPClient()

	prot := protocol.NewGeneric(phy)

	for {
		var f frame.Frame
		var pk packet.Packet
		prot.FromNetworkLayer(&pk)
		fmt.Println("FromNetworkLayer", pk)

		f.Info = pk

		prot.ToPhysicalLayer(&f)
		fmt.Println("ToPhysicalLayer", f)
	}
}

func (u *utopianSimplex) Stop() {
	u.done <- true
}

func (u *utopianSimplex) Wait() bool {
	return <-u.done
}
