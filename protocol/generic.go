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
	buffer       []byte
}

func NewGeneric(p physical.Interface) Interface {
	return &generic{
		physical: p,
		event:    make(chan event.Event),
		buffer:   []byte{},
	}
}

func (g *generic) WaitForEvent() <-chan event.Event {
	go func() {
		for {
			b, err := g.physical.Recv()
			if err != nil {
				g.event <- event.Event{Type: event.ERROR}
				continue
			}

			g.buffer = b
			g.event <- event.Event{Type: event.FRAME_ARRIVAL}
		}
	}()

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
	if len(g.buffer) == 0 {
		panic("Buffer vacío")
	}

	nf, err := frame.FromBytes(g.buffer)
	if err != nil {
		panic(err)
	}

	*f = *nf

	g.buffer = []byte{}
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
