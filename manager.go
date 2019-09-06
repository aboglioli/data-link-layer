package main

type Manager interface {
	ConvertToFrames([]byte) ([]*Frame, error)
	ConvertToBytes([]*Frame) ([]byte, error)
}

type manager struct {
	serializer Serializer
}

func NewManager() Manager {
	return &manager{
		serializer: NewSerializer(),
	}
}

func (m *manager) ConvertToFrames(b []byte) ([]*Frame, error) {
	frame, err := m.serializer.BytesToFrame(b)
	if err != nil {
		return nil, err
	}

	return []*Frame{
		frame,
	}, nil
}

func (m *manager) ConvertToBytes(f []*Frame) ([]byte, error) {
	bytes, err := m.serializer.FrameToBytes(f[0])
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
