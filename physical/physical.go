package physical

// Interfaz de comunicación
type Transmissor interface {
	ToPhysicalLayer([]byte) error
	FromPhysicalLayer() ([]byte, error)
}
