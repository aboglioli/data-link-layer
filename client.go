package main

import "net"

type Client struct {
	transmissor Transmissor
	manager     Manager
}

func ConnectClient() (*Client, error) {
	c := GetConfig()
	conn, err := net.Dial(c.Communication, c.Address())
	if err != nil {
		return nil, err
	}

	return &Client{
		transmissor: NewTCPTransmissor(conn),
		manager:     NewManager(),
	}, nil
}

func NewClient(t Transmissor) *Client {
	return &Client{
		transmissor: t,
		manager:     NewManager(),
	}
}

func (c *Client) Send(f *Frame) error {
	msg, err := c.manager.ConvertToBytes([]*Frame{f})
	if err != nil {
		return err
	}

	return c.transmissor.ToPhysicalLayer(msg)
}

func (c *Client) Recv() (*Frame, error) {
	msg, err := c.transmissor.FromPhysicalLayer()
	if err != nil {
		return nil, err
	}

	f, err := c.manager.ConvertToFrames(msg)

	return f[0], err
}
