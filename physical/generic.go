package physical

import (
	"errors"
	"io"
)

type generic struct {
	rw io.ReadWriter
}

func NewTransmissor(rw io.ReadWriter) Interface {
	return &generic{rw}
}

func (g *generic) ToPhysicalLayer(msg []byte) error {
	l, err := g.rw.Write(msg)
	if l <= 0 || err != nil {
		return errors.New("Error en envío")
	}

	return nil
}

func (g *generic) FromPhysicalLayer() ([]byte, error) {
	msg := make([]byte, 4096)
	l, err := g.rw.Read(msg)
	if err != nil || l <= 0 {
		return nil, errors.New("Error en recepción")
	}

	return msg[0:l], nil
}
