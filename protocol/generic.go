package protocol

import (
	"fmt"
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
	fmt.Println("[FromNetworkLayer:start]")

	b := <-net.Sender

	fmt.Println("[FromNetworkLayer:middle]")

	np, err := packet.FromBytes(b)
	if err != nil {
		panic(err)
	}

	*p = *np

	fmt.Println("[FromNetworkLayer:end]")
}

func (g *generic) ToNetworkLayer(p *packet.Packet) {
	net := network.Get()

	fmt.Println("[ToNetworkLayer:start]")
	b, err := p.ToBytes()
	if err != nil {
		panic(err)
	}

	fmt.Println("[ToNetworkLayer:middle]")

	net.Receiver <- b

	fmt.Println("[ToNetworkLayer:end]")
}

func (g *generic) FromPhysicalLayer(f *frame.Frame) {
	fmt.Println("[FromPhysicalLayer:start]")
	b, err := g.physical.Recv()
	if err != nil {
		g.event <- event.Event{Type: event.ERROR}
	}

	fmt.Println("[FromPhysicalLayer:middle]")

	nf, err := frame.FromBytes(b)
	if err != nil {
		panic(err)
	}

	*f = *nf

	fmt.Println("[FromPhysicalLayer:end]")
}

func (g *generic) ToPhysicalLayer(f *frame.Frame) {
	fmt.Println("[ToPhysicalLayer:start]")
	b, err := f.ToBytes()
	if err != nil {
		panic(err)
	}

	fmt.Println("[ToPhysicalLayer:middle]")

	err = g.physical.Send(b)
	if err != nil {
		panic(err)
	}

	fmt.Println("[ToPhysicalLayer:end]")
}

func (g *generic) StartTimer()          {}
func (g *generic) StopTimer()           {}
func (g *generic) EnableNetworkLayer()  {}
func (g *generic) DisableNetworkLayer() {}
