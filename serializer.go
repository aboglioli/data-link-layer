package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Serializer interface {
	FrameToString(*Frame) (string, error)
	StringToFrame(string) (*Frame, error)
}

type serializer struct {
	sep string
}

func NewSerializer() Serializer {
	return &serializer{
		sep: ":", // separador
	}
}

func (s *serializer) FrameToString(f *Frame) (string, error) {
	if f.Ack < 0 || f.Seq < 0 {
		return "", errors.New("SEQ o ACK inválido")
	}

	if f.Payload == "" {
		return "", errors.New("Mensaje vacío")
	}

	return fmt.Sprintf("%d:%d:%s", f.Seq, f.Ack, f.Payload), nil
}

func (s *serializer) StringToFrame(str string) (*Frame, error) {
	fmt.Println(str)
	if strings.Count(str, ":") != 2 {
		return nil, errors.New("Trama inválida")
	}

	arr := strings.Split(str, ":")
	if len(arr) != 3 {
		return nil, errors.New("Trama inválida")
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
