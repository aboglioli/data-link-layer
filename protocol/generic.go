package protocol

import (
	"io"

	"github.com/aboglioli/data-link-layer/event"
	"github.com/aboglioli/data-link-layer/frame"
	"github.com/aboglioli/data-link-layer/network"
	"github.com/aboglioli/data-link-layer/packet"
	"github.com/aboglioli/data-link-layer/physical"
)

type generic struct {
	physical     physical.Interface
	event        chan event.Event // TODO: not used yet
	networkLayer io.ReadWriter
}

func NewGeneric(p physical.Interface) Interface {
	return &generic{
		physical: p,
		event:    make(chan event.Event),
	}
}

func (g *generic) WaitForEvent() <-chan event.Event {
	return g.event
}

func (g *generic) FromNetworkLayer(p *packet.Packet) {
	net := network.Get()

	b := <-net.Sender

	np, err := packet.FromBytes(b)
	if err != nil {
		panic(err)
	}

	*p = *np
}

func (g *generic) ToNetworkLayer(p *packet.Packet) {
	net := network.Get()

	b := p.ToBytes()

	net.Receiver <- b
}

func (g *generic) FromPhysicalLayer(f *frame.Frame) {
	b, err := g.physical.Recv()
	if err != nil {
		g.event <- event.Event{Type: event.ERROR}
	}

	nf, err := frame.FromBytes(b)
	if err != nil {
		panic(err)
	}

	*f = *nf
}

func (g *generic) ToPhysicalLayer(f *frame.Frame) {
	b := f.ToBytes()

	err := g.physical.Send(b)
	if err != nil {
		panic(err)
	}
}

func (g *generic) StartTimer()          {}
func (g *generic) StopTimer()           {}
func (g *generic) EnableNetworkLayer()  {}
func (g *generic) DisableNetworkLayer() {}
