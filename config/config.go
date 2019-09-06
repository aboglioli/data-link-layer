package config

import (
	"fmt"
	"sync"
)

const (
	SERIAL = "serial"
	TCP    = "tcp"
	UDP    = "udp"
)

type Config struct {
	Communication  string
	Port           int
	Host           string
	MaxFrames      int
	MinFrameLength int
	MaxFrameLength int
}

func (c *Config) Address() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

var (
	c *Config
	o sync.Once
)

func Get() *Config {
	o.Do(func() {
		c = &Config{
			Communication:  TCP,
			Port:           7788,
			MaxFrames:      3,
			MinFrameLength: 5,
			MaxFrameLength: 64,
		}
	})

	return c
}
