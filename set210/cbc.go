package set210

import (
	"crypto/aes"
	"encoding/hex"

	"github.com/shrayolacrayon/cryptopals/set12"
	"github.com/shrayolacrayon/cryptopals/set29"
	"github.com/shrayolacrayon/cryptopals/util"
)

func EncryptBlock(block []byte, key []byte) ([]byte, error) {
	cipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	destination := make([]byte, len(block))
	cipher.Encrypt(destination, block)
	return destination, nil
}

func DecryptBlock(block []byte, key []byte) ([]byte, error) {
	cipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	dest := make([]byte, len(block))
	cipher.Decrypt(dest, block)
	return dest, nil
}

func CBCFile(inputFilepath string, key []byte) ([]byte, error) {
	allBytes, err := util.ReadFile(inputFilepath)
	if err != nil {
		return nil, err
	}
	return CBC(allBytes, key)
}

// with an IV of all ASCII 0 (\x00\x00\x00 &c)
// TODO: what is this?

func CBC(allBytes []byte, key []byte) ([]byte, error) {

	paddedAllBytes := set29.PKCS7(allBytes, len(key))

	blocks := util.CreateBlocks(paddedAllBytes, len(key))

	// the IV is the fake first block
	mainBlock := make([]byte, len(key))
	for i := range mainBlock {
		mainBlock[i] = byte(0)
	}
	allBlocks := []byte{}
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
		allBlocks = append(allBlocks, combinedBlock...)
	}
	return allBlocks, nil

}

// with an IV of all ASCII 0 (\x00\x00\x00 &c)
// TODO: what is this?

func CBCAnotherTry(allBytes []byte, key []byte) ([]byte, error) {

	paddedAllBytes := set29.PKCS7(allBytes, len(key))

	blocks := util.CreateBlocks(paddedAllBytes, len(key))

	// the IV is the fake first block
	mainBlock := make([]byte, len(key))
	for i := range mainBlock {
		mainBlock[i] = byte(0)
	}
	allBlocks := []byte{}
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
		allBlocks = append(allBlocks, combinedBlock...)
	}
	return allBlocks, nil

}

func CBCOtherDecrypt(input []byte, key []byte) ([]byte, error) {
	allBlocks := []byte{}
	iv := make([]byte, len(key))
	for i := range iv {
		iv[i] = byte(0)
	}
	blocks := append([][]byte{iv}, util.CreateBlocks(input, len(key))...)
	for i := 1; i < len(blocks); i++ {
		// xor with the encrypted block
		xorBlock, err := set12.FixedXOR(hex.EncodeToString(blocks[i]), hex.EncodeToString(blocks[i-1]))
		if err != nil {
			return nil, err
		}
		decrypted, err := DecryptBlock(xorBlock, key)
		if err != nil {
			return nil, err
		}
		allBlocks = append(allBlocks, decrypted...)
	}

	return allBlocks, nil
}

func CBCDecryptFile(inputFilepath string, key []byte) ([]byte, error) {
	allBytes, err := util.ReadFile(inputFilepath)
	if err != nil {
		return nil, err
	}
	return CBCOtherDecrypt(allBytes, key)
}

func CBCDecrypt(input []byte, key []byte) ([]byte, error) {
	allBlocks := []byte{}
	iv := make([]byte, len(key))
	for i := range iv {
		iv[i] = byte(0)
	}
	// needs to be xor'ed again
	blocks := util.CreateBlocks(input, len(key))
	// start with the last block
	mainBlock := blocks[len(blocks)-1]
	for i := len(blocks) - 2; i >= 0; i-- {
		//for i := range blocks {
		// each block is appended with the iv
		xorBlock, err := set12.FixedXOR(hex.EncodeToString(mainBlock), hex.EncodeToString(blocks[i]))
		if err != nil {
			return nil, err
		}

		decrypted, err := DecryptBlock(xorBlock, key)
		if err != nil {
			return nil, err
		}

		mainBlock = xorBlock
		allBlocks = append(decrypted, allBlocks...)
	}

	xorBlock, err := set12.FixedXOR(hex.EncodeToString(mainBlock), hex.EncodeToString(iv))
	decrypted, err := DecryptBlock(xorBlock, key)
	if err != nil {
		return nil, err
	}
	allBlocks = append(decrypted, allBlocks...)
	return allBlocks, nil
}
