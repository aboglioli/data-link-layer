package main

import (
	"fmt"

	"github.com/aboglioli/data-link-layer/implementation"
	"github.com/aboglioli/data-link-layer/network"
)

func receiver() {
	us := implementation.UtopianSimplex()
	go us.StartReceiver()

	net := network.Get()

	for b := range net.Receiver {
		fmt.Println("Recibido:", b, " -> ", string(b))
	}

	us.Wait()
}

func main() {
	receiver()
}
