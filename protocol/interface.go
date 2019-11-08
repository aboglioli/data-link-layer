package protocol

import (
	"github.com/aboglioli/data-link-layer/event"
	"github.com/aboglioli/data-link-layer/frame"
	"github.com/aboglioli/data-link-layer/packet"
)

// Interface a implementar por el protocolo de la capa de enlace de datos.
// Idéntica a la presentada en Redes de Información de Tanenbaum.
type Interface interface {
	// Comunicación asíncrona, espera eventos
	WaitForEvent() <-chan event.Event

	// Comunicación con la capa de Red
	FromNetworkLayer(*packet.Packet)
	ToNetworkLayer(*packet.Packet)

	// Comunicación con la capa física
	FromPhysicalLayer(*frame.Frame)
	ToPhysicalLayer(*frame.Frame)

	// Timers
	StartTimer()
	StopTimer()

	// Permite que la capa de red pueda interrumpir a la capa de enlace para enviar
	// más paquetes, o no.
	EnableNetworkLayer()
	DisableNetworkLayer()
}
