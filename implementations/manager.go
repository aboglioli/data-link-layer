package implementations

import (
	"github.com/aboglioli/data-link-layer/conversors"
	"github.com/aboglioli/data-link-layer/interfaces"
	"github.com/aboglioli/data-link-layer/types"
)

type manager struct{}

func NewManager() interfaces.Protocol {
	return &manager{}
}

func (m *manager) ConvertToFrames(b []byte) (types.Frames, error) {
	frame, err := conversors.BytesToFrame(b)
	if err != nil {
		return nil, err
	}

	return types.Frames{
		frame,
	}, nil
}

func (m *manager) ConvertToBytes(f types.Frames) ([]byte, error) {
	bytes, err := conversors.FrameToBytes(f[0])
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
