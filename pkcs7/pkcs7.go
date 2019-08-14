package pkcs7

// PKCS7 adds PKCS7 padding to a byte block
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

// PKCS7Decrypt takes in an input and strips off the PKCS7 padding, erroring with invalid padding
func PKCS7Decrypt(input []byte) ([]byte, error) {
	// get the last byte of the input, which would have the padding
	end := input[len(input)-1]
	// check if end is a number from 0 to 255

}
