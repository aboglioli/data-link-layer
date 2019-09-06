package main

import "net"

type Client struct {
	transmissor Transmissor
}

func NewClient() (*Client, error) {
	c := GetConfig()
	conn, err := net.Dial(c.Communication, c.Address())
	if err != nil {
		return nil, err
	}

	return &Client{
		transmissor: NewTCPTransmissor(conn),
	}, nil
}

func (c *Client) Send(f *Frame) error {
	return c.transmissor.ToPhysicalLayer(f)
}

func (c *Client) Recv() (*Frame, error) {
	return c.transmissor.FromPhysicalLayer()
}
