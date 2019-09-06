package implementations

import (
	"errors"

	"github.com/aboglioli/data-link-layer/config"
	"github.com/aboglioli/data-link-layer/types"
)

type protocol struct {
	payloadLength    int
	frames           types.Frames
	concurrentFrames int
}

func NewProtocol() *protocol {
	c := config.Get()
	return &protocol{
		payloadLength:    c.PayloadLength,
		concurrentFrames: c.ConcurrentFrames,
	}
}

func (p *protocol) PrepareMessage(msg string) (types.Frames, error) {
	if msg == "" {
		return nil, errors.New("Mensaje vacío")
	}

	payloads := make([]string, 0)
	for i := 0; i < len(msg); i += p.payloadLength {
		payloads = append(payloads, msg[i:(i+p.payloadLength)])
	}

	frames := make(types.Frames, 0)
	for i, p := range payloads {
		frames = append(frames, types.NewFrame(i, 0, p))
	}

	p.frames = frames

	return frames, nil
}
