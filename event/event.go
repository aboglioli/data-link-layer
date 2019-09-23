package event

type EventType string

const (
	ERROR          EventType = "error"
	FRAME_ARRIVAL            = "frame_arrival"
	FRAME_RECEIVED           = "frame_received"
)

type Event struct {
	Type EventType
}
