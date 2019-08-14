package util

import (
	"encoding/base64"
	"io/ioutil"
	"os"
)

func ReadFile(filepath string) ([]byte, error) {
	// read in the whole file
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	allBytes, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return allBytes, nil
}

func DecryptBase64File(filepath string) ([]byte, error) {
	allBytes, err := ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	allBytes, err = base64.StdEncoding.DecodeString(string(allBytes))
	if err != nil {
		return nil, err
	}
	return allBytes, nil
}

func CreateBlocks(allBytes []byte, size int) [][]byte {
	blocks := [][]byte{}
	end := len(allBytes) / size

	for i := 0; i < end; i++ {
		blocks = append(blocks, allBytes[i*size:(i+1)*size])
	}
	if len(allBytes) > (size * end) {
		blocks = append(blocks, allBytes[size*end:])
	}
	return blocks
}

func SumBlockCounts(data []byte) float64 {
	blockCounts := map[string]int{}
	// figure out if the line is decoded by an ecb
	// it can only be a base4 so 16 or 32 bytes, repeated in some way
	// split up the encrypted into 16 bytes and see if something is repeated
	keysize := 16
	blocks := CreateBlocks(data, keysize)
	for _, block := range blocks {
		blockCounts[string(block)]++
	}
	sum := 0
	// sum the squares?
	for _, count := range blockCounts {
		sum += count * count
	}
	return float64(sum) / float64(len(blocks))
}
