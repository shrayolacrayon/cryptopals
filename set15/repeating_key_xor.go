package set15

import (
	"encoding/hex"

	"github.com/shrayolacrayon/cryptopals/set12"
)

func RepeatingKeyXOR(input, xor string) []byte {
	encoded := hex.EncodeToString([]byte(input))
	// convert string to byte array
	single := []byte(xor)
	xorbytes := []byte{}
	for i := 0; i < len(input)/(len(single))+1; i++ {
		xorbytes = append(xorbytes, single...)
	}
	encodedXOR := hex.EncodeToString(xorbytes[:len(input)])
	val, err := set12.FixedXOR(encoded, encodedXOR)
	if err != nil {
		panic(err)
	}
	return val
}
