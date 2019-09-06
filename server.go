package main

import (
	"fmt"
	"net"

	"github.com/aboglioli/data-link-layer/config"
)

type Server struct {
	listener net.Listener
}

func NewServer() (*Server, error) {
	c := config.Get()
	listener, err := net.Listen(c.Communication, c.Address())
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

			c <- NewClient(NewTCPTransmissor(conn))
		}
	}()

	return c
}
