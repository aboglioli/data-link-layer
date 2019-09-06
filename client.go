package main

import "net"

type Client struct {
	Transmissor Transmissor
}

func NewClient() (*Client, error) {
	c := GetConfig()
	conn, err := net.Dial(c.Communication, c.Address())
	if err != nil {
		return nil, err
	}

	return &Client{
		Transmissor: NewTCPTransmissor(conn),
	}, nil
}
