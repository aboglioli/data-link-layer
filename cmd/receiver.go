package main

import (
	"fmt"

	"github.com/aboglioli/data-link-layer/event"
	"github.com/aboglioli/data-link-layer/frame"
	"github.com/aboglioli/data-link-layer/physical"
	"github.com/aboglioli/data-link-layer/protocol"
)

func receiver() {
	fmt.Println("Receptor iniciado")

	s := physical.TCPServer()
	fmt.Println("Interfaz física")

	p := protocol.UtopianSimplex(s)
	fmt.Println("Implementación del protocolo")

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

func main() {
	receiver()
}
