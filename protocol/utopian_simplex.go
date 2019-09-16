package protocol

import (
	"fmt"
	"io"

	"github.com/aboglioli/data-link-layer/event"
	"github.com/aboglioli/data-link-layer/frame"
	"github.com/aboglioli/data-link-layer/packet"
	"github.com/aboglioli/data-link-layer/physical"
)

type utopianSimplex struct {
	physical     physical.Interface
	event        chan event.Event
	networkLayer io.ReadWriter
}

func UtopianSimplex(p physical.Interface) Interface {
	return &utopianSimplex{
		physical: p,
		event:    make(chan event.Event),
	}
}

func (u *utopianSimplex) WaitForEvent() <-chan event.Event {
	return u.event
}

func (u *utopianSimplex) FromNetworkLayer(p *packet.Packet) {
	*p = packet.Packet{Data: "empty"}
}

func (u *utopianSimplex) ToNetworkLayer(p *packet.Packet) {
	fmt.Println("To network layer", p)
}

func (u *utopianSimplex) FromPhysicalLayer(f *frame.Frame) {
	_, err := u.physical.Recv()
	if err != nil {
		u.event <- event.Event{Type: event.ERROR}
	}

	*f = frame.Frame{
		Kind: frame.DATA,
		Seq:  0,
		Ack:  0,
		Info: packet.Packet{
			Data: "empty",
		},
	}
}

func (u *utopianSimplex) ToPhysicalLayer(f *frame.Frame) {
	err := u.physical.Send([]byte("empty"))
	if err != nil {
		u.event <- event.Event{Type: event.ERROR}
	}
}

func (u *utopianSimplex) StartTimer()          {}
func (u *utopianSimplex) StopTimer()           {}
func (u *utopianSimplex) EnableNetworkLayer()  {}
func (u *utopianSimplex) DisableNetworkLayer() {}
