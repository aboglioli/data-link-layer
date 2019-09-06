package main

import (
	"flag"
	"strings"
)

func startServer() {
	s, err := NewServer()
	if err != nil {
		panic(err)
	}

	s.Listen()
}

func startClient() {
	c, err := NewClient()
	if err != nil {
		panic(err)
	}

	c.Transmissor.ToPhysicalLayer(&Frame{12, 13, "Hola"})
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
