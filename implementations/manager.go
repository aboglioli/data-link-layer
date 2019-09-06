package implementations

import (
	"github.com/aboglioli/data-link-layer/converters"
	"github.com/aboglioli/data-link-layer/frame"
	"github.com/aboglioli/data-link-layer/interfaces"
)

type manager struct{}

func NewManager() interfaces.Protocol {
	return &manager{}
}

func (m *manager) ConvertToFrames(b []byte) (frame.Frames, error) {
	f, err := converters.BytesToFrame(b)
	if err != nil {
		return nil, err
	}

	return frame.Frames{
		f,
	}, nil
}

func (m *manager) ConvertToBytes(f frame.Frames) ([]byte, error) {
	bytes, err := converters.FrameToBytes(f[0])
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
