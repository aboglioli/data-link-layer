package frame

import (
	"errors"
	"strconv"
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

func FromRaw(r *RawFrame) (*Frame, error) {
	seq, err := strconv.Atoi(string(r.Seq))
	if err != nil {
		return nil, err
	}

	ack, err := strconv.Atoi(string(r.Ack))
	if err != nil {
		return nil, err
	}

	payload := string(r.Payload[:])
	checksum := string(r.Checksum[:])

	if payload != checksum {
		return nil, errors.New("Diferencia entre payload y checksum")
	}

	flag := getFlagFromByte(r.Flags)

	return New(flag, seq, ack, payload), nil
}

func (f *Frame) NextSeq() {
	f.Seq++
}

func (f *Frame) NextAck() {
	f.Ack++
}

func getFlagFromByte(b byte) Flag {
	switch b {
	case 0xF0:
		return SYN
	case 0x0F:
		return FIN
	case 0xFF:
		return ACK
	}

	return ERROR
}
