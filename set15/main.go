package main

import (
	"encoding/hex"

	"github.com/shrayolacrayon/cryptopals/set12"
)

var stanza = "Burning 'em, if you ain't quick and nimble \n I go crazy when I hear a cymbal"

func repeatingKeyXOR(input, xor string) {
	encoded, err := hex.EncodeToString([]byte(input))
	if err != nil {
		panic(err)
	}
	// convert string to byte array
	xorbytes := []byte(xor)

	// create a repeating string with the xor string the length of input
	xorbytes := make([]byte, len(encoded))

}
