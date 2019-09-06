package protocol

type Interface interface {
	// Detección de errores
	CheckError([]byte) bool

	// Servicios de transmisón
	// utiliza frames (tramas) de fondo
	Send([]byte) error
	Recv() ([]byte, error)
}
