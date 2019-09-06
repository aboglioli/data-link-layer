package main

import (
	"errors"
	"net"

	"github.com/aboglioli/data-link-layer/interfaces"
)

// Implementación sobre TCP y Sockets
type connection struct {
	socket net.Conn
}

func NewTCPTransmissor(conn net.Conn) interfaces.Transmissor {
	return &connection{
		socket: conn,
	}
}

func (c *connection) ToPhysicalLayer(msg []byte) error {
	l, err := c.socket.Write(msg)
	if l <= 0 || err != nil {
		return errors.New("Error en envío")
	}

	return nil
}

func (c *connection) FromPhysicalLayer() ([]byte, error) {
	msg := make([]byte, 4096)
	l, err := c.socket.Read(msg)
	if err != nil || l <= 0 {
		return nil, errors.New("Error en recepción")
	}

	return msg[0:l], nil
}
