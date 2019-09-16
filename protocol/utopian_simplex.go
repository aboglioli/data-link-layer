package protocol

import (
	"github.com/aboglioli/data-link-layer/event"
	"github.com/aboglioli/data-link-layer/physical"
)

type utopianSimplex struct {
	transmissor physical.Interface
	event       chan event.Event
}

func UtopianSimplex(t physical.Interface) Interface {
	return &utopianSimplex{
		transmissor: t,
		event:       make(chan event.Event),
	}
}
