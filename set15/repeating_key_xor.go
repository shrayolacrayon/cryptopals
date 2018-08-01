package set15

import (
	"encoding/hex"

	"github.com/shrayolacrayon/cryptopals/set12"
)

var stanza = "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
var expected = "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"

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
