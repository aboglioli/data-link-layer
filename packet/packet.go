package packet

type Packet struct {
	Data string
}

func New(d string) *Packet {
	return &Packet{
		Data: d,
	}
}

func FromBytes(b []byte) (*Packet, error) {
	return New(string(b)), nil
}

func (p *Packet) ToBytes() ([]byte, error) {
	return []byte(p.Data), nil
}
