package set11

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func HexToBase64(inputHex string) (string, error) {
	decoded, err := hex.DecodeString(inputHex)
	if err != nil {
		return "", err
	}
	fmt.Printf("%s \n", decoded)
	return base64.StdEncoding.EncodeToString(decoded), nil
}
