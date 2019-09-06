package main

import (
	"net"

	"github.com/aboglioli/data-link-layer/config"
	"github.com/aboglioli/data-link-layer/frame"
	"github.com/aboglioli/data-link-layer/physical"
	"github.com/aboglioli/data-link-layer/protocol"
)

type Client struct {
	transmissor physical.Transmissor
	protocol    protocol.Interface
}

func ConnectClient() (*Client, error) {
	c := config.Get()
	conn, err := net.Dial(c.CommunicationMethod(), c.Address())
	if err != nil {
		return nil, err
	}

	return &Client{
		transmissor: physical.NewTCPTransmissor(conn),
	}, nil
}

func NewClient(t physical.Transmissor) *Client {
	return &Client{
		transmissor: t,
	}
}

func (c *Client) Send(f *frame.Frame) error {
	msg, err := f.ToBytes()
	if err != nil {
		return err
	}

	return c.transmissor.ToPhysicalLayer(msg)
}

func (c *Client) Recv() (*frame.Frame, error) {
	msg, err := c.transmissor.FromPhysicalLayer()
	if err != nil {
		return nil, err
	}

	return frame.FromBytes(msg)
}
