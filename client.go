package main

import (
	"net"

	"github.com/aboglioli/data-link-layer/config"
	"github.com/aboglioli/data-link-layer/implementations"
	"github.com/aboglioli/data-link-layer/interfaces"
	"github.com/aboglioli/data-link-layer/types"
)

type Client struct {
	transmissor interfaces.Transmissor
	protocol    interfaces.Protocol
}

func ConnectClient() (*Client, error) {
	c := config.Get()
	conn, err := net.Dial(c.CommunicationMethod(), c.Address())
	if err != nil {
		return nil, err
	}

	return &Client{
		transmissor: implementations.NewTCPTransmissor(conn),
		protocol:    implementations.NewManager(),
	}, nil
}

func NewClient(t interfaces.Transmissor) *Client {
	return &Client{
		transmissor: t,
		protocol:    implementations.NewManager(),
	}
}

func (c *Client) Send(f *types.Frame) error {
	msg, err := c.protocol.ConvertToBytes(types.Frames{f})
	if err != nil {
		return err
	}

	return c.transmissor.ToPhysicalLayer(msg)
}

func (c *Client) Recv() (*types.Frame, error) {
	msg, err := c.transmissor.FromPhysicalLayer()
	if err != nil {
		return nil, err
	}

	f, err := c.protocol.ConvertToFrames(msg)

	return f[0], err
}
