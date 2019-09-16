package implementation

import (
	"fmt"

	"github.com/aboglioli/data-link-layer/event"
	"github.com/aboglioli/data-link-layer/frame"
	"github.com/aboglioli/data-link-layer/packet"
	"github.com/aboglioli/data-link-layer/physical"
	"github.com/aboglioli/data-link-layer/protocol"
)

type UtopianSimplex struct {
}

func UtopianSimplex() *UtopianSimplex {
	return &UtopianSimplex{}
}

func (u *UtopianSimplex) StartReceiver() {
	fmt.Println("Receptor iniciado")

	s := physical.TCPServer()
	fmt.Println("Esperando clientes")

	p := protocol.UtopianSimplex(s)

	fmt.Println("Esperando por eventos")
	for e := range p.WaitForEvent() {
		if e.Type == event.ERROR || e.Type == event.FRAME_ARRIVAL {
			fmt.Println("[ERROR]", e)
		}

		var f frame.Frame

		p.FromPhysicalLayer(&f)
		fmt.Println("FromPhysicalLayer", f)

		p.ToNetworkLayer(&f.Info)
		fmt.Println("ToNetworkLayer", f.Info)
	}
}

func (u *UtopianSimplex) StartSender() {
	fmt.Println("Emisor iniciado")

	s := physical.TCPClient()
	fmt.Println("Esperando servidor")

	p := protocol.UtopianSimplex(s)

	for {
		var f frame.Frame
		var pk packet.Packet
		p.FromNetworkLayer(&pk)
		fmt.Println("FromNetworkLayer", pk)

		f.Info = pk

		p.ToPhysicalLayer(&f)
		fmt.Println("ToPhysicalLayer", f)
	}
}
