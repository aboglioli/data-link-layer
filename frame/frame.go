package frame

import (
	"github.com/aboglioli/data-link-layer/config"
	"github.com/aboglioli/data-link-layer/packet"
)

// Trama
type SeqNr uint

type FrameKind string

const (
	DATA FrameKind = "DATA"
	ACK            = "ACK"
	NAK            = "NAK"
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
		Kind: k,
		Seq:  seq,
		Ack:  ack,
		Info: info,
	}
}

func FromBytes(b []byte) (*Frame, error) {
	p, err := packet.FromBytes(b)
	if err != nil {
		return nil, err
	}

	// TODO: implement real frame
	return New(DATA, 0, 0, *p), nil
}

func (f *Frame) ToBytes() ([]byte, error) {
	return []byte("empty"), nil
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
