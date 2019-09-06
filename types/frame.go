package types

type Frame struct {
	Seq     int
	Ack     int
	Payload string
}

type Frames []*Frame

func NewFrame(s int, a int, p string) *Frame {
	return &Frame{
		Seq:     s,
		Ack:     a,
		Payload: p,
	}
}

func (f *Frame) NextSeq() {
	f.Seq++
}

func (f *Frame) NextAck() {
	f.Ack++
}
