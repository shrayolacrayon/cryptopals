package set11

import (
	"encoding/base64"
	"encoding/hex"
)

func HexToBase64(inputHex string) string {
	decoded, err := hex.DecodeString(inputHex)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(decoded)
}
