package config

import (
	"fmt"
	"sync"
	"time"
)

type Channel int

const (
	SERIAL Channel = iota
	TCP
	UDP
)

type Config struct {
	Communication    Channel
	Port             int
	Host             string
	ConcurrentFrames int
	MinFrameLength   int
	MaxFrameLength   int
	Timeout          time.Duration
	PayloadLength    int
}

func (c *Config) Address() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

func (c *Config) CommunicationMethod() string {
	switch c.Communication {
	case SERIAL:
		return "serial"
	case TCP:
		return "tcp"
	case UDP:
		return "udp"
	default:
		return "tcp"
	}
}

var (
	c *Config
	o sync.Once
)

func Get() *Config {
	o.Do(func() {
		c = &Config{
			Communication:    TCP,
			Port:             7788,
			ConcurrentFrames: 3,
			MinFrameLength:   5,
			MaxFrameLength:   64,
			Timeout:          500 * time.Millisecond,
			PayloadLength:    3,
		}
	})

	return c
}
