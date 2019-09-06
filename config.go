package main

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
	Communication string
	Port          int
	Host          string
}

func (c *Config) Address() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

var (
	c *Config
	o sync.Once
)

func GetConfig() *Config {
	o.Do(func() {
		c = &Config{
			Communication: TCP,
			Port:          7788,
		}
	})

	return c
}
