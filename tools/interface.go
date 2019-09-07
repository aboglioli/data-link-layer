package tools

type Transmissor interface {
	Send(string) error
	Recv() (string, error)
}

type Receiver interface {
	Listen() Sender
}

type Sender interface {
	Connect() Transmissor
}
