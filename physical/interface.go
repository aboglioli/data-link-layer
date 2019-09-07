package physical

// Interfaz de comunicación
type Interface interface {
	ToPhysicalLayer([]byte) error
	FromPhysicalLayer() ([]byte, error)
}
