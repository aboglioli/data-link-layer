package protocol

import (
	"fmt"
	"testing"
)

func TestSplitMessgeIntoFrames(t *testing.T) {
	p := NewProtocol()

	// División en tramas
	f, err := p.PrepareMessage("Esto es una prueba")
	if err != nil {
		t.Error(err)
	}

	if len(f) != 8 {
		t.Errorf("Se esperan 8 tramas, se obtuvieron %d\n", len(f))
	}

	// Error por mensaje vacío
	_, err = p.PrepareMessage("")
	fmt.Println(err)

	if err.Error() != "Mensaje vacío" {
		t.Error("Se espera error por mensaje vacío")
	}
}

func TestGenerateSeq(t *testing.T) {
	p := NewProtocol()

	// División en tramas
	f, err := p.PrepareMessage("Esto es una prueba")
	if err != nil {
		t.Error(err)
	}

	if len(f) != 8 {
		t.Errorf("Se esperan 8 tramas, se obtuvieron %d\n", len(f))
	}
}
