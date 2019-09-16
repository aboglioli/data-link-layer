package frame

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/aboglioli/data-link-layer/config"
	"github.com/aboglioli/data-link-layer/packet"
)

// Trama
type SeqNr uint

type FrameKind string

const (
	DATA FrameKind = "DATA"
	ACK  FrameKind = "ACK"
	NAK  FrameKind = "NAK"
)

type Frame struct {
	Kind FrameKind
	Seq  SeqNr
	Ack  SeqNr
	Info packet.Packet
}

type Frames []*Frame

func New(k FrameKind, seq SeqNr, ack SeqNr, info packet.Packet) *Frame {
	return &Frame{
		k,
		seq,
		ack,
		info,
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

	t := FrameKind(arr[0])

	seq, err := strconv.Atoi(arr[1])
	if err != nil {
		return nil, errors.New("SEQ inválido")
	}

	ack, err := strconv.Atoi(arr[2])
	if err != nil {
		return nil, errors.New("ACK inválido")
	}

	payload := packet.Packet{arr[2]}

	return New(t, SeqNr(seq), SeqNr(ack), payload), nil
}

func (f *Frame) NextSeq() SeqNr {
	c := config.Get()

	if f.Seq < SeqNr(c.MaxSeq) {
		f.Seq++
	} else {
		f.Seq = 0
	}

	return f.Seq
}

func (f *Frame) NextAck() SeqNr {
	c := config.Get()

	if f.Ack < SeqNr(c.MaxSeq) {
		f.Ack++
	} else {
		f.Ack = 0
	}

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

	str := fmt.Sprintf("%s:%d:%d:%s", f.Kind, f.Seq, f.Ack, f.Info.Data)

	c := config.Get()
	if len(str) < c.MinFrameLength || len(str) > c.MaxFrameLength {
		return nil, errors.New("Tamaño de trama")
	}

	return []byte(str), nil
}
