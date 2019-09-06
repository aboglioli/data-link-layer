package frame

type RawFrame struct {
	Seq      byte
	Ack      byte
	Payload  [4]byte
	Flags    byte // NS, CWR, ECE, URG, ACK, PSH, RST, SYN, FIN
	Checksum [4]byte
}
