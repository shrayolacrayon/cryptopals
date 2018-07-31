package set12

import (
	"encoding/hex"
	"fmt"
)

func FixedXOR(input1, input2 string) ([]byte, error) {
	decode1, err := hex.DecodeString(input1)
	if err != nil {
		return nil, err
	}
	decode2, err := hex.DecodeString(input2)
	if err != nil {
		return nil, err
	}
	if len(decode1) != len(decode2) {
		return nil, fmt.Errorf("not the same length strings")
	}

	newBytes := make([]byte, len(decode1))
	for i, i1 := range decode1 {
		newBytes[i] = i1 ^ decode2[i]
	}

	return newBytes, nil
}
