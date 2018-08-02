package set18

import (
	"encoding/hex"
	"strings"

	"github.com/shrayolacrayon/cryptopals/util"
)

// Remember that the problem with ECB is that it is stateless and deterministic;
// the same 16 byte plaintext block will always produce the same 16 byte ciphertext.
func DecryptFiles(filepath string) ([]byte, error) {
	hexBytes, err := util.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	maxSum := 0
	var maxline []byte
	lines := strings.Split(string(hexBytes), "\n")
	for _, line := range lines {
		decoded, err := hex.DecodeString(line)
		if err != nil {
			return nil, err
		}
		blockCounts := map[string]int{}
		// figure out if the line is decoded by an ecb
		// it can only be a base4 so 16 or 32 bytes, repeated in some way
		// split up the encrypted into 16 bytes and see if something is repeated?
		keysize := 16
		blocks := util.CreateBlocks(decoded, keysize)
		for _, block := range blocks {
			blockCounts[string(block)]++
		}
		sum := 0
		// sum the squares?
		for _, count := range blockCounts {
			sum += count * count
		}
		if sum > maxSum {
			maxline = decoded
			maxSum = sum
		}
	}
	return maxline, nil
}
