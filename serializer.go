package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Serializer interface {
	FrameToBytes(*Frame) ([]byte, error)
	BytesToFrame([]byte) (*Frame, error)
}

type serializer struct {
	sep string
}

func NewSerializer() Serializer {
	return &serializer{
		sep: ":", // separador
	}
}

func (s *serializer) FrameToBytes(f *Frame) ([]byte, error) {
	if f.Ack < 0 || f.Seq < 0 {
		return nil, errors.New("SEQ o ACK inválido")
	}

	if f.Payload == "" {
		return nil, errors.New("Mensaje vacío")
	}

	str := fmt.Sprintf("%d:%d:%s", f.Seq, f.Ack, f.Payload)

	c := GetConfig()
	if len(str) < c.MinFrameLength || len(str) > c.MaxFrameLength {
		return nil, errors.New("Tamaño de trama")
	}

	return []byte(str), nil
}

func (s *serializer) BytesToFrame(bytes []byte) (*Frame, error) {
	str := string(filterMessage(bytes))

	c := GetConfig()
	if len(str) < c.MinFrameLength || len(str) > c.MaxFrameLength {
		return nil, errors.New("Tamaño de trama")
	}

	if strings.Count(str, ":") != 2 {
		return nil, errors.New("Separadores de trama inválidos")
	}

	arr := strings.Split(str, ":")
	if len(arr) != 3 {
		return nil, errors.New("Separadores de trama inválidos")
	}

	seq, err := strconv.Atoi(arr[0])
	if err != nil {
		return nil, errors.New("SEQ inválido")
	}

	ack, err := strconv.Atoi(arr[1])
	if err != nil {
		return nil, errors.New("ACK inválido")
	}

	payload := arr[2]

	if len(payload) <= 0 {
		return nil, errors.New("Payload vacío")
	}

	return NewFrame(seq, ack, payload), nil
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
