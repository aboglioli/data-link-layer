package main

import (
	"fmt"
	"time"

	"github.com/aboglioli/data-link-layer/implementation"
	"github.com/aboglioli/data-link-layer/network"
)

func receiver() {
	us := implementation.UtopianSimplex()
	go us.StartReceiver()
	time.Sleep(1 * time.Second)

	net := network.Get()

	b := <-net.Receiver
	fmt.Println(b)

	us.Wait()
}

func main() {
	receiver()
}
