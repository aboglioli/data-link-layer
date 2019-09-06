package main

type Frame struct {
	Seq     int
	Ack     int
	Payload string
}

func NewFrame(s int, a int, p string) *Frame {
	return &Frame{
		Seq:     s,
		Ack:     a,
		Payload: p,
	}
}
