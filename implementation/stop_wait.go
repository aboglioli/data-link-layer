package implementation

import (
	"github.com/aboglioli/data-link-layer/frame"
	"github.com/aboglioli/data-link-layer/packet"
	"github.com/aboglioli/data-link-layer/physical"
	"github.com/aboglioli/data-link-layer/protocol"
)

type stopWait struct {
	done chan bool
}

func StopWait() *stopWait {
	return &stopWait{
		done: make(chan bool),
	}
}

func (s *stopWait) StartReceiver() {
	phy := physical.TCPServer()

	prot := protocol.NewGeneric(phy)

	for {
		var f frame.Frame
		prot.FromPhysicalLayer(&f)
		prot.ToNetworkLayer(&f.Info)
	}
}

func (s *stopWait) StartSender() {
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

func (s *stopWait) Stop() {
	s.done <- true
}

func (s *stopWait) Wait() bool {
	return <-s.done
}
