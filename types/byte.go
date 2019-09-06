package types

import "fmt"

type ByteCount struct {
	ones   int
	zeroes int
}

// Un byte
type Byte byte

func (b Byte) String() string {
	return fmt.Sprintf("%b", b)
}

func (b Byte) Count() ByteCount {
	var ones, zeroes int

	for _, s := range b.String() {
		c := string(s)

		if c == "1" {
			ones++
		} else {
			zeroes++
		}
	}

	return ByteCount{ones, zeroes}
}

// Múltiples byes
type Bytes []Byte

func (b Bytes) Concat(bytes ...Byte) Bytes {
	return append(b, bytes...)
}

func (b Bytes) String() string {
	return fmt.Sprintf("%b", b)
}

func (b Bytes) Count() ByteCount {
	count := ByteCount{0, 0}

	for _, byte := range b {
		bc := byte.Count()
		count.ones = count.ones + bc.ones
		count.zeroes = count.zeroes + bc.zeroes
	}

	return count
}
