package physical

import (
	"net"
)

// Implementación sobre TCP y Sockets
func NewTCPTransmissor(rw net.Conn) Interface {
	return &generic{rw}
}
