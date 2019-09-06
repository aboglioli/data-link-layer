package main

import (
	"encoding/binary"
	"testing"
	"time"
)

func TestBinary(t *testing.T) {
	buf := make([]byte, 10)
	ts := uint32(time.Now().Unix())
	binary.BigEndian.PutUint16(buf[0:], 0xa20c) // sensorID
	binary.BigEndian.PutUint16(buf[2:], 0x04af) // locationID
	binary.BigEndian.PutUint32(buf[4:], ts)     // timestamp
	binary.BigEndian.PutUint16(buf[8:], 479)    // temp
	t.Logf("% x\n", buf)
}
