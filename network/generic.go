package network

// La implementación de la capa de red es básica, consta de dos canales: uno
// para recepción, otro para emisión. Estos simulan los buffers utilizados
// realmente para la comunicación entre capas.
type Generic struct {
	Sender   chan []byte
	Receiver chan []byte
}

var generic *Generic

func Get() *Generic {
	if generic == nil {
		generic = &Generic{
			Sender:   make(chan []byte),
			Receiver: make(chan []byte),
		}
	}

	return generic
}
