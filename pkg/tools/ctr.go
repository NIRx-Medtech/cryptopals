package tools

type KeyStream interface {
	NextByte() byte
}

func CtrStream(input []byte, ks KeyStream) []byte {
	xored := make([]byte, len(input))
	for i, b := range input {
		xored[i] = b ^ ks.NextByte()
	}
	return xored
}
