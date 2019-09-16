package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/aboglioli/data-link-layer/frame"
)

func startServer() {
	s, err := NewServer()
	if err != nil {
		panic(err)
	}

	client := s.Listen()
	for c := range client {
		f, err := c.Recv()
		if err != nil {
			fmt.Println("[ERROR]", err)
		} else {
			fmt.Println(f)
			c.Send(&frame.Frame{frame.ACK, 1, 1, frame.Packet{"OK"}})
		}
	}
}

func startClient() {
	c, err := ConnectClient()
	if err != nil {
		panic(err)
	}

	err = c.Send(&frame.Frame{frame.ACK, 12, 13, frame.Packet{"Hola"}})
	if err != nil {
		panic(err)
	}
	f, _ := c.Recv()
	fmt.Println(f)
}

func main() {
	flagMode := flag.String("mode", "server", "correr Client o Servidor")
	flag.Parse()

	if strings.ToLower(*flagMode) == "server" {
		startServer()
	} else {
		startClient()
	}
}
