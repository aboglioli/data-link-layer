package interfaces

import (
	"github.com/aboglioli/data-link-layer/types"
)

type Manager interface {
	ConvertToFrames([]byte) (types.Frames, error)
	ConvertToBytes(types.Frames) ([]byte, error)
}
