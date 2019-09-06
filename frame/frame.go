package frame

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/aboglioli/data-link-layer/config"
)

// Flags: SYN, FIN, ACK, étc
type Flag string

const (
	SYN   Flag = "SYN"   // 11110000 = 0xF0
	FIN        = "FIN"   // 00001111 = 0x0F
	ACK        = "ACK"   // 11111111 = 0xFF
	ERROR      = "ERROR" // 00000000 = 0x00
)

// Trama
type Frame struct {
	Type    Flag
	Seq     int
	Ack     int
	Payload string
}

type Frames []*Frame

func New(f Flag, seq int, ack int, payload string) *Frame {
	return &Frame{
		f,
		seq,
		ack,
		payload,
	}
}

func FromBytes(bytes []byte) (*Frame, error) {
	str := string(filterMessage(bytes))

	c := config.Get()
	if len(str) < c.MinFrameLength || len(str) > c.MaxFrameLength {
		return nil, errors.New("Tamaño de trama")
	}

	if strings.Count(str, ":") != 3 {
		return nil, errors.New("Separadores de trama inválidos")
	}

	arr := strings.Split(str, ":")
	if len(arr) != 3 {
		return nil, errors.New("Separadores de trama inválidos")
	}

	t := Flag(arr[0])

	seq, err := strconv.Atoi(arr[1])
	if err != nil {
		return nil, errors.New("SEQ inválido")
	}

	ack, err := strconv.Atoi(arr[2])
	if err != nil {
		return nil, errors.New("ACK inválido")
	}

	payload := arr[2]

	return New(t, seq, ack, payload), nil
}

func (f *Frame) NextSeq() int {
	f.Seq++
	return f.Seq
}

func (f *Frame) NextAck() int {
	f.Ack++
	return f.Ack
}

func (f *Frame) Swap() {
	ack := f.Ack
	f.Ack = f.Seq + 1
	f.Seq = ack
}

func (f *Frame) ToBytes() ([]byte, error) {
	if f.Ack < 0 || f.Seq < 0 {
		return nil, errors.New("SEQ o ACK inválido")
	}

	str := fmt.Sprintf("%s:%d:%d:%s", f.Type, f.Seq, f.Ack, f.Payload)

	c := config.Get()
	if len(str) < c.MinFrameLength || len(str) > c.MaxFrameLength {
		return nil, errors.New("Tamaño de trama")
	}

	return []byte(str), nil
}
