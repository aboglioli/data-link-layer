package conversors

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/aboglioli/data-link-layer/config"
	"github.com/aboglioli/data-link-layer/types"
)

func FrameToBytes(f *types.Frame) ([]byte, error) {
	if f.Ack < 0 || f.Seq < 0 {
		return nil, errors.New("SEQ o ACK inválido")
	}

	if f.Payload == "" {
		return nil, errors.New("Mensaje vacío")
	}

	str := fmt.Sprintf("%d:%d:%s", f.Seq, f.Ack, f.Payload)

	c := config.Get()
	if len(str) < c.MinFrameLength || len(str) > c.MaxFrameLength {
		return nil, errors.New("Tamaño de trama")
	}

	return []byte(str), nil
}

func BytesToFrame(bytes []byte) (*types.Frame, error) {
	str := string(filterMessage(bytes))

	c := config.Get()
	if len(str) < c.MinFrameLength || len(str) > c.MaxFrameLength {
		return nil, errors.New("Tamaño de trama")
	}

	if strings.Count(str, ":") != 2 {
		return nil, errors.New("Separadores de trama inválidos")
	}

	arr := strings.Split(str, ":")
	if len(arr) != 3 {
		return nil, errors.New("Separadores de trama inválidos")
	}

	seq, err := strconv.Atoi(arr[0])
	if err != nil {
		return nil, errors.New("SEQ inválido")
	}

	ack, err := strconv.Atoi(arr[1])
	if err != nil {
		return nil, errors.New("ACK inválido")
	}

	payload := arr[2]

	if len(payload) <= 0 {
		return nil, errors.New("Payload vacío")
	}

	return types.NewFrame(seq, ack, payload), nil
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
