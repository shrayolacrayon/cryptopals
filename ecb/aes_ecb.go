package ecb

// Also what is needed for set17

import (
	"crypto/aes"

	"github.com/shrayolacrayon/cryptopals/util"
)

var KEYSIZE = 16

func Encrypt(allBytes []byte, key []byte) ([]byte, error) {
	cipher, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}
	allEncrypted := []byte{}
	blocks := util.CreateBlocks(allBytes, KEYSIZE)
	for _, block := range blocks {
		encrypted := make([]byte, KEYSIZE)
		cipher.Encrypt(encrypted, block)
		if err != nil {
			return allEncrypted, err
		}
		allEncrypted = append(allEncrypted, encrypted...)
	}
	return allEncrypted, nil
}

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
