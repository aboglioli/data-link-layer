package implementation

import (
	"fmt"

	"github.com/aboglioli/data-link-layer/frame"
	"github.com/aboglioli/data-link-layer/packet"
	"github.com/aboglioli/data-link-layer/physical"
	"github.com/aboglioli/data-link-layer/protocol"
)

// Aquí se implementa un protocolo de capa de enlace utópica simplex: un emisor
// envía y un receptor recibe tramas simples.

// Esta implementación consiste en un receptor y un emisor. Ambos deben correr
// en paralelo, por eso se utilizan los hilos de go para correrlos de fondo y
// que no interfieran con el flujo normal de la aplicación. Se utilizan canales
// para simular la comunicación a través de buffers que existen en las
// implementaciones reales de los protocolos.
type utopianSimplex struct {
	done chan bool
}

func UtopianSimplex() *utopianSimplex {
	return &utopianSimplex{
		done: make(chan bool),
	}
}

// Receptor
// - Recibe eventos y reacciones
// - Toma bytes de la capa física y los convierte a tramas.
// - Direcciona dichas tramas a la capa de red. Utiliza canales para dicha comunicación.
func (u *utopianSimplex) StartReceiver() {
	phy := physical.TCPServer()

	prot := protocol.NewGeneric(phy)

	// Se espera por
	for e := range prot.WaitForEvent() {
		fmt.Println("Evento:", e)

		var f frame.Frame
		prot.FromPhysicalLayer(&f)
		prot.ToNetworkLayer(&f.Info)
	}
}

// Emisor
// - Toma paquetes de la capa de red, los transforma a tramas. Carga útil del
// paquete es insertada en una trama.
// - Envía tramas a la capa física para que esta pueda enviarlas a través del
// canal de comunicación.
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

// Como esta capa corre en background, es necesario saber cuándo está
// ejecutándose para que la aplicación no se detenga repentinamente.
func (u *utopianSimplex) Wait() bool {
	return <-u.done
}
