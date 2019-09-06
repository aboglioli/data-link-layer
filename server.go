package main

import (
	"errors"
	"fmt"
	"net"
)

type Server struct {
	listener      net.Listener
	transmissor   Transmissor
	NewConnection chan bool
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
			fmt.Println("[ERROR]", err)
		}

		s.transmissor = NewTCPTransmissor(conn)
		f, err := s.Recv()
		fmt.Println(f)
	}
}

func (s *Server) Send(f *Frame) error {
	if s.transmissor != nil {
		return s.transmissor.ToPhysicalLayer(f)
	}

	return errors.New("No hay clientes conecteados")
}

func (s *Server) Recv() (*Frame, error) {
	if s.transmissor != nil {
		return s.transmissor.FromPhysicalLayer()
	}

	return nil, errors.New("No hay clientes conectados")
}
