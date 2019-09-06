package frame

func filterMessage(msg []byte) []byte {
	filter := make([]byte, 0, len(msg))
	for _, b := range msg {
		if b != 0 {
			filter = append(filter, b)
		}
	}
	return filter
}
