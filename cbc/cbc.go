package cbc

import (
	"crypto/aes"
	"encoding/hex"

	"github.com/shrayolacrayon/cryptopals/pkcs7"
	"github.com/shrayolacrayon/cryptopals/set12"
	"github.com/shrayolacrayon/cryptopals/util"
)

func AESEncryptBlock(block []byte, key []byte) ([]byte, error) {
	cipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	destination := make([]byte, len(block))
	cipher.Encrypt(destination, block)
	return destination, nil
}

func AESDecryptBlock(block []byte, key []byte) ([]byte, error) {
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
	return CBCEncrypt(allBytes, key)
}

func CBCEncrypt(allBytes []byte, key []byte) ([]byte, error) {

	paddedAllBytes := pkcs7.PKCS7(allBytes, len(key))

	blocks := util.CreateBlocks(paddedAllBytes, len(key))

	// the IV is the fake first block
	mainBlock := make([]byte, len(key))
	for i := range mainBlock {
		mainBlock[i] = byte(0)
	}
	allBlocks := []byte{}
	for _, block := range blocks {

		// first xor the blocks
		combinedBlock, err := set12.FixedXOR(hex.EncodeToString(mainBlock), hex.EncodeToString(block))
		if err != nil {
			return nil, err
		}

		encrypted, err := AESEncryptBlock(combinedBlock, key)
		if err != nil {
			return nil, err
		}

		mainBlock = encrypted
		allBlocks = append(allBlocks, encrypted...)
	}
	return allBlocks, nil

}

func CBCDecrypt(input []byte, key []byte) ([]byte, error) {
	allBlocks := []byte{}
	iv := make([]byte, len(key))
	for i := range iv {
		iv[i] = byte(0)
	}
	blocks := append([][]byte{iv}, util.CreateBlocks(input, len(key))...)
	for i := 1; i < len(blocks); i++ {
		decrypted, err := AESDecryptBlock(blocks[i], key)
		if err != nil {
			return nil, err
		}

		// xor with the encrypted block
		xorBlock, err := set12.FixedXOR(hex.EncodeToString(decrypted), hex.EncodeToString(blocks[i-1]))
		if err != nil {
			return nil, err
		}
		allBlocks = append(allBlocks, xorBlock...)
	}

	return allBlocks, nil
}
