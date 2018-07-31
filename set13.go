package main

import (
	"github.com/shrayolacrayon/cryptopals/set13"
)

func main() {
	// set13
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	xorVal, output, _ := set13.XORChar(input)
	fmt.Printf("%#U : %s", xorVal, output)
}
