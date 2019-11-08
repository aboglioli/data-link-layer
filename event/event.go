package event

// Los posibles eventos de una trama para una implementación básica: Se producen
// errores, llega una nueva trama desde la capa física o se recibe una nueva
// trama desde la capa de red.
type EventType string

const (
	ERROR          EventType = "error"
	FRAME_ARRIVAL            = "frame_arrival"
	FRAME_RECEIVED           = "frame_received"
)

type Event struct {
	Type EventType
}
