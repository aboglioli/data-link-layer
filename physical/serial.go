package physical

import (
	"github.com/tarm/serial"
)

// Implementación sobre Puerto serial
func NewSerialTransmissor() Interface {
	c := &serial.Config{Name: "COM5", Baud: 115200}
	port, err := serial.OpenPort(c)
	if err != nil {
		panic(err)
	}

	return &generic{
		rw: port,
	}
}
