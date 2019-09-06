package converters

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/aboglioli/data-link-layer/config"
	"github.com/aboglioli/data-link-layer/frame"
)

func FrameToBytes(f *frame.Frame) ([]byte, error) {
	if f.Ack < 0 || f.Seq < 0 {
		return nil, errors.New("SEQ o ACK inválido")
	}

	str := fmt.Sprintf("%s:%d:%d:%s", f.Type, f.Seq, f.Ack, f.Payload)

	c := config.Get()
	if len(str) < c.MinFrameLength || len(str) > c.MaxFrameLength {
		return nil, errors.New("Tamaño de trama")
	}

	return []byte(str), nil
}

func BytesToFrame(bytes []byte) (*frame.Frame, error) {
	str := string(filterMessage(bytes))

	c := config.Get()
	if len(str) < c.MinFrameLength || len(str) > c.MaxFrameLength {
		return nil, errors.New("Tamaño de trama")
	}

	if strings.Count(str, ":") != 3 {
		return nil, errors.New("Separadores de trama inválidos")
	}

	arr := strings.Split(str, ":")
	if len(arr) != 3 {
		return nil, errors.New("Separadores de trama inválidos")
	}

	t := frame.Flag(arr[0])

	seq, err := strconv.Atoi(arr[1])
	if err != nil {
		return nil, errors.New("SEQ inválido")
	}

	ack, err := strconv.Atoi(arr[2])
	if err != nil {
		return nil, errors.New("ACK inválido")
	}

	payload := arr[2]

	return frame.New(t, seq, ack, payload), nil
}

func filterMessage(msg []byte) []byte {
	filter := make([]byte, 0, len(msg))
	for _, b := range msg {
		if b != 0 {
			filter = append(filter, b)
		}
	}
	return filter
}
