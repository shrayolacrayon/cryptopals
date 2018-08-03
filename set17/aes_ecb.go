package set17

import (
	"crypto/aes"

	"github.com/shrayolacrayon/cryptopals/util"
)

var key = "YELLOW SUBMARINE"
var KEYSIZE = 16

func Decrypt(filepath string, key []byte) []byte {
	allBytes, err := util.DecryptBase64File(filepath)
	if err != nil {
		panic(err)
	}
	cipher, err := aes.NewCipher([]byte(key))

	allDecrypted := []byte{}
	chunked := util.CreateBlocks(allBytes, KEYSIZE)
	for _, chunks := range chunked {
		decrypted := make([]byte, KEYSIZE)
		cipher.Decrypt(decrypted, chunks)
		if err != nil {
			panic(err)
		}
		allDecrypted = append(allDecrypted, decrypted...)
	}
	return allDecrypted

}
