package implementations

import (
	"errors"

	"github.com/aboglioli/data-link-layer/interfaces"
	"github.com/tarm/serial"
)

// Implementación sobre TCP y Sockets
type serialTransmissor struct {
	port *serial.Port
}

func NewSerialTransmissor() interfaces.Transmissor {
	c := &serial.Config{Name: "COM5", Baud: 115200}
	port, err := serial.OpenPort(c)
	if err != nil {
		panic(err)
	}

	return &serialTransmissor{
		port: port,
	}
}

func (c *serialTransmissor) ToPhysicalLayer(msg []byte) error {
	l, err := c.port.Write(msg)
	if l <= 0 || err != nil {
		return errors.New("Error en envío")
	}

	return nil
}

func (c *serialTransmissor) FromPhysicalLayer() ([]byte, error) {
	msg := make([]byte, 4096)
	l, err := c.port.Read(msg)
	if err != nil || l <= 0 {
		return nil, errors.New("Error en recepción")
	}

	return msg[0:l], nil
}
