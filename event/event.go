package event

type EventType string

const (
	FRAME_ARRIVAL EventType = "frame_arrival"
)

type Event struct {
	Type EventType
}
