package main

import (
	"fmt"

	"github.com/shrayolacrayon/cryptopals/set11"
)

func main() {
	input := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	output, err := set11.HexToBase64(input)
	if err != nil {
		fmt.Printf("error: %s", err)
	}
	fmt.Printf("output: %s", output)
}
