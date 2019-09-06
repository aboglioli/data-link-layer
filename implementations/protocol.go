package implementations

import (
	"errors"

	"github.com/aboglioli/data-link-layer/config"
	"github.com/aboglioli/data-link-layer/frame"
	"github.com/aboglioli/data-link-layer/interfaces"
)

type protocol struct {
	payloadLength int
	transmissor   interfaces.Transmissor
}

func NewProtocol() *protocol {
	c := config.Get()
	return &protocol{
		payloadLength: c.PayloadLength,
	}
}

func (p *protocol) PrepareMessage(msg string) (frame.Frames, error) {
	if msg == "" {
		return nil, errors.New("Mensaje vacío")
	}

	payloads := make([]string, 0)
	for i := 0; i < len(msg); i += p.payloadLength {
		payloads = append(payloads, msg[i:(i+p.payloadLength)])
	}

	frames := make(frame.Frames, 0)

	frames = append(frames, frame.New(frame.SYN, 0, 0, ""))
	for _, p := range payloads {
		frames = append(frames, frame.New(frame.ACK, 0, 0, p))
	}
	frames = append(frames, frame.New(frame.FIN, 7, 0, ""))

	return frames, nil
}
