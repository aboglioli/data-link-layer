package event

type EventType string

const (
	ERROR         EventType = "error"
	FRAME_ARRIVAL           = "frame_arrival"
)

type Event struct {
	Type EventType
}
