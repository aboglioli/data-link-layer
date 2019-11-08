package main

import (
	"fmt"

	"github.com/aboglioli/data-link-layer/implementation"
	"github.com/aboglioli/data-link-layer/network"
)

func receiver() {
	// Inicializa la implementación
	us := implementation.UtopianSimplex()

	// Inicializa el receptor en un hilo corriendo en background
	go us.StartReceiver()

	// E inicializa la capa de red. Esta, simplemente simula la capa de red al
	// emitir y recibir mensajes a través de canales. Gracias a Go podemos
	// utilizar canales a nivel de lenguaje, los cuales emularían el uso de
	// buffers a través de los cuales se comunican capa de red y de enlace.
	net := network.Get()

	// De fondo, la implementación del receptor toma bytes de la capa física,
	// los convierte a trama y los pasa a la capa de red. Aquí podemos apreciar
	// cómo al capa de red percibe dichas tramas y cómo las recibe. El
	// funcionamiento de la capa de enlace no es visible aquí ya que así debe
	// ser. La capa de red no sabe nada respecto a este, excepto el servicio que presta.
	// Ver implementación de Utopian Simplex.
	for b := range net.Receiver {
		fmt.Println("Recibido:", b, " -> ", string(b))
	}

	us.Wait()
}

func main() {
	receiver()
}
