package main

import (
	"fmt"

	"github.com/aboglioli/data-link-layer/frame"
	"github.com/aboglioli/data-link-layer/packet"
	"github.com/aboglioli/data-link-layer/physical"
	"github.com/aboglioli/data-link-layer/protocol"
)

func sender() {
	fmt.Println("Emisor iniciado")

	s := physical.TCPClient()
	fmt.Println("Interfaz física")

	p := protocol.UtopianSimplex(s)
	fmt.Println("Implementación del protocolo")

	for {
		var f frame.Frame
		var pk packet.Packet
		p.FromNetworkLayer(&pk)
		f.Info = pk
		p.ToPhysicalLayer(&f)
	}
}

func main() {
	sender()
}
