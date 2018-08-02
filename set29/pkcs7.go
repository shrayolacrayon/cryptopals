package set29

func PKCS7(input []byte, blockSize int) []byte {
	toAdd := blockSize - (len(input) % blockSize)
	if !(toAdd > 0 && toAdd < 255) {
		panic("difference for padding is too large")
	}
	additionalBytes := make([]byte, toAdd)
	for i := range additionalBytes {
		additionalBytes[i] = byte(toAdd)
	}
	return append(input, additionalBytes...)

}
