package main

import (
	"fmt"
	"net"

	"github.com/aboglioli/data-link-layer/config"
	"github.com/aboglioli/data-link-layer/physical"
)

type Server struct {
	listener net.Listener
}

func NewServer() (*Server, error) {
	c := config.Get()
	listener, err := net.Listen(c.CommunicationMethod(), c.Address())
	if err != nil {
		return nil, err
	}

	return &Server{
		listener: listener,
	}, nil
}

func (s *Server) Listen() <-chan *Client {
	c := make(chan *Client)

	go func() {
		for {
			conn, err := s.listener.Accept()
			if err != nil {
				fmt.Println("[ERROR]", err)
			}

			c <- NewClient(physical.NewTransmissor(conn))
		}
	}()

	return c
}
