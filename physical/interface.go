package physical

// Interface de comunicación.
// Servicios prestados por la capa física.
// Se envían bytes desde las capas físicas de emisor y receptor.
type Interface interface {
	Send([]byte) error
	Recv() ([]byte, error)
}
