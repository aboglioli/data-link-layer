package physical

import (
	"net"
)

// Implementación sobre TCP y Sockets
// Simple simulación utilizando sockets: un servidor <-> un cliente.
func TCPServer() Interface {
	ln, err := net.Listen("tcp", ":7788")
	if err != nil {
		panic(err)
	}

	conn, err := ln.Accept()
	if err != nil {
		panic(err)
	}

	return NewTransmissor(conn)
}

func TCPClient() Interface {
	conn, err := net.Dial("tcp", "localhost:7788")
	if err != nil {
		panic(err)
	}

	return NewTransmissor(conn)
}
