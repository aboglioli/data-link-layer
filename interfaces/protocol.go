package interfaces

import (
	"github.com/aboglioli/data-link-layer/frame"
)

type Protocol interface {
	ConvertToFrames([]byte) (frame.Frames, error)
	ConvertToBytes(frame.Frames) ([]byte, error)
}
