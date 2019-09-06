package implementations

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

	if len(f) != 6 {
		t.Errorf("Se esperan 6 tramas, se obtuvieron %d\n", len(f))
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

	if len(f) != 6 {
		t.Errorf("Se esperan 6 tramas, se obtuvieron %d\n", len(f))
	}

	for i := 0; i < len(f); i++ {
		if f[i].Seq != i {
			t.Errorf("Número de seq %d inválido para trama %d", i, f[i].Seq)
		}
	}
}
