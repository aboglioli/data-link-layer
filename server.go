package main

import (
	"fmt"
	"net"
)

type Server struct {
	listener    net.Listener
	Transmissor Transmissor
}

func NewServer() (*Server, error) {
	c := GetConfig()
	listener, err := net.Listen(c.Communication, c.Address())
	if err != nil {
		return nil, err
	}

	return &Server{
		listener: listener,
	}, nil
}

func (s *Server) Listen() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			fmt.Println(err)
		}

		s.Transmissor = NewTCPTransmissor(conn)
		f, err := s.Transmissor.FromPhysicalLayer()
		fmt.Println(f, err)
	}
}
