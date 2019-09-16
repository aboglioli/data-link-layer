package network

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
