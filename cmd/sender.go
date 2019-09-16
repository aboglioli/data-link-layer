package main

import (
	"fmt"
	"time"

	"github.com/aboglioli/data-link-layer/implementation"
	"github.com/aboglioli/data-link-layer/network"
)

func sender() {
	us := implementation.UtopianSimplex()
	go us.StartSender()
	time.Sleep(2 * time.Second)

	net := network.Get()
	fmt.Println("Sending...")
	net.Sender <- []byte("Hola")
	fmt.Println("Sent")

	us.Wait()
}

func main() {
	sender()
}
