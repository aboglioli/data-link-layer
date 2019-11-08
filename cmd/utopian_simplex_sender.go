package main

import (
	"fmt"

	"github.com/aboglioli/data-link-layer/implementation"
	"github.com/aboglioli/data-link-layer/network"
)

func sender() {
	us := implementation.UtopianSimplex()

	// Levanta emisor en otro hilo para que no interfiera con el principal.
	go us.StartSender()

	net := network.Get()

	// Se simula la capa de red intentando emitir una nueva trama.
	// De fondo, la implementación del emisor tomará esta trama de esta capa de
	// red y la pasará a la capa física. Aquí no vemos nada de eso por ser una
	// abstracción.
	fmt.Println("Enviando: \"Hola\"")
	net.Sender <- []byte("Hola")

	fmt.Println("Enviando: \"Chau\"")
	net.Sender <- []byte("Chau")

	us.Wait()
}

func main() {
	sender()
}
