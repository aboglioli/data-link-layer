package main

import (
	"fmt"
	"net"
)

type Server struct {
	listener net.Listener
	manager  Manager
}

func NewServer() (*Server, error) {
	c := GetConfig()
	listener, err := net.Listen(c.Communication, c.Address())
	if err != nil {
		return nil, err
	}

	return &Server{
		listener: listener,
		manager:  NewManager(),
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
