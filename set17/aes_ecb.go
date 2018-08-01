package set17

import (
	"crypto/aes"

	"github.com/shrayolacrayon/cryptopals/set16"
)

var key = "YELLOW SUBMARINE"
var KEYSIZE = 16

func Decrypt(filepath string) []byte {
	allBytes := set16.DecryptBase64(filepath)
	cipher, err := aes.NewCipher([]byte(key))

	allDecrypted := []byte{}
	for i := 0; i < len(allBytes)/KEYSIZE; i++ {
		decrypted := make([]byte, KEYSIZE)
		cipher.Decrypt(decrypted, allBytes[i*KEYSIZE:(i+1)*KEYSIZE])
		if err != nil {
			panic(err)
		}
		allDecrypted = append(allDecrypted, decrypted...)
	}
	return allDecrypted

}
