package network

type Interface interface {
	Send([]byte) error
	Recv() ([]byte, error)

	ToSend() []byte
	ToRecv([]byte)
}
