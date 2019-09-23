package main

import (
	"fmt"

	"github.com/aboglioli/data-link-layer/implementation"
	"github.com/aboglioli/data-link-layer/network"
)

func sender() {
	us := implementation.UtopianSimplex()
	go us.StartSender()

	net := network.Get()

	fmt.Println("Enviando: \"Hola\"")
	net.Sender <- []byte("Hola")

	fmt.Println("Enviando: \"Chau\"")
	net.Sender <- []byte("Chau")

	us.Wait()
}

func main() {
	sender()
}
