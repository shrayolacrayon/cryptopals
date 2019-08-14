package set211

import (
	"crypto/rand"
	"fmt"
	random "math/rand"

	"github.com/shrayolacrayon/cryptopals/cbc"
	"github.com/shrayolacrayon/cryptopals/ecb"
	"github.com/shrayolacrayon/cryptopals/util"
)

type BlockType int

const (
	ECB BlockType = iota
	CBC
)

// GenerateKey creates a random set of bytes of length size
func GenerateKey(size int) ([]byte, error) {
	b := make([]byte, size)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func encrypt(data []byte, mode int) ([]byte, error) {
	var encrypted []byte

	key, err := GenerateKey(16)
	if err != nil {
		return encrypted, err
	}
	switch BlockType(mode) {
	case ECB:
		encrypted, err = ecb.Encrypt(data, key)
	case CBC:
		encrypted, err = cbc.CBCEncrypt(data, key)
	default:
		return nil, fmt.Errorf("options for modes are ECB and CBC")
	}
	return encrypted, err
}

func EncryptionOracle(data []byte) ([]byte, error) {
	padding := random.Intn(5) + 5
	newBytes := make([]byte, len(data)+2*padding)
	additionalBytes := make([]byte, padding)
	for i := range additionalBytes {
		additionalBytes[i] = byte(padding)
	}
	newBytes = append(newBytes, additionalBytes...)
	newBytes = append(newBytes, data...)
	newBytes = append(newBytes, additionalBytes...)

	return encrypt(newBytes, random.Intn(1))
}

// What should this threshold be?
var threshold = 1.0

// Detect the block cipher mode the function is using each time.
// You should end up with a piece of code that, pointed at a block box that might be encrypting ECB or CBC, tells you which one is happening.
func DetectMethod(encrypted []byte) BlockType {
	// we know that ecb is detected via looking for blocks so we can use the same method that we used in set18
	score := util.SumBlockCounts(encrypted)
	if score > threshold {
		return ECB
	}
	return CBC
}
