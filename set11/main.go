package set11

import (
	"encoding/base64"
	"encoding/hex"
	"log"
)

func hexToBase64(inputHex string) (string, error) {
	decoded, err := hex.DecodeString(inputHex)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(decoded), nil
}

func main() {
	input := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	output, err := hexToBase64(input)
	if err != nil {
		log.Printf("error: %s", err)
	}
	log.Printf("output: %s", output)
}
