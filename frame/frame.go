package frame

import (
	"encoding/binary"

	"github.com/aboglioli/data-link-layer/config"
	"github.com/aboglioli/data-link-layer/packet"
)

// Número de secuencia
type SeqNr uint32

func (s SeqNr) ToBytes() []byte {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, uint32(s))
	return b
}

// Tipo de frame
type FrameKind string

func (f FrameKind) ToBytes() []byte {
	return []byte(f)
}

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

func (f *Frame) ToBytes() []byte {
	var b []byte
	b = append(b, f.Kind.ToBytes()...)
	b = append(b, f.Seq.ToBytes()...)
	b = append(b, f.Ack.ToBytes()...)
	b = append(b, f.Info.ToBytes()...)
	return b
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
