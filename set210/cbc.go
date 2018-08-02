package set210

import (
	"crypto/aes"
	"encoding/hex"

	"github.com/shrayolacrayon/cryptopals/set12"
	"github.com/shrayolacrayon/cryptopals/set29"
	"github.com/shrayolacrayon/cryptopals/util"
)

func EncryptBlock(allBytes []byte, key []byte) ([]byte, error) {
	cipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	destination := make([]byte, len(allBytes))
	cipher.Encrypt(destination, allBytes)
	if err != nil {
		return nil, err
	}
	return destination, nil

}

// with an IV of all ASCII 0 (\x00\x00\x00 &c)
// TODO: what is this?

func CBC(inputFilepath string, key []byte) ([]byte, error) {
	allBytes, err := util.ReadFile(inputFilepath)
	if err != nil {
		return nil, err
	}

	paddedAllBytes := set29.PKCS7(allBytes, len(key))

	blocks := util.CreateBlocks(paddedAllBytes, len(key))

	// the IV is the fake first block
	mainBlock := make([]byte, len(key))
	for i := range mainBlock {
		mainBlock[i] = byte(0)
	}

	for _, block := range blocks {
		encrypted, err := EncryptBlock(block, key)
		if err != nil {
			return nil, err
		}

		// each block is appended with the iv
		combinedBlock, err := set12.FixedXOR(hex.EncodeToString(mainBlock), hex.EncodeToString(encrypted))
		if err != nil {
			return nil, err
		}
		mainBlock = combinedBlock
	}
	return mainBlock, err

}
