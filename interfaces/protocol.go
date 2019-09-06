package interfaces

import (
	"github.com/aboglioli/data-link-layer/types"
)

type Protocol interface {
	ConvertToFrames([]byte) (types.Frames, error)
	ConvertToBytes(types.Frames) ([]byte, error)
}
