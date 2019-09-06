package main

import (
	"errors"
	"net"
)

// Interfaz de comunicación
type Transmissor interface {
	ToPhysicalLayer(f *Frame) error
	FromPhysicalLayer() (*Frame, error)
}

// Implementación sobre TCP y Sockets
type connection struct {
	socket     net.Conn
	serializer Serializer
}

func NewTCPTransmissor(conn net.Conn) Transmissor {
	return &connection{
		socket:     conn,
		serializer: NewSerializer(),
	}
}

func (c *connection) ToPhysicalLayer(f *Frame) error {
	str, err := c.serializer.FrameToString(f)
	if err != nil {
		return err
	}

	l, err := c.socket.Write([]byte(str))
	if l <= 0 || err != nil {
		return errors.New("Error en envío")
	}

	return nil
}

func (c *connection) FromPhysicalLayer() (*Frame, error) {
	msg := make([]byte, 4096)
	l, err := c.socket.Read(msg)
	if err != nil || l <= 0 {
		return nil, errors.New("Error en recepción")
	}

	return c.serializer.StringToFrame(string(filterMessage(msg)))
}

func filterMessage(msg []byte) []byte {
	filter := make([]byte, 0, len(msg))
	for _, b := range msg {
		if b != 0 {
			filter = append(filter, b)
		}
	}
	return filter
}
