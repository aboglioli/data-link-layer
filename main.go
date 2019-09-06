package main

import (
	"flag"
	"fmt"
	"strings"
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
		}
	}
}

func startClient() {
	c, err := ConnectClient()
	if err != nil {
		panic(err)
	}

	err = c.Send(&Frame{12, 13, "Hola"})
	if err != nil {
		panic(err)
	}
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
