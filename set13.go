package main

import (
	"github.com/shrayolacrayon/cryptopals/set13"
)

func main() {
	// set13
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	training := set13.CreateTrainingMap("testdata/11-0.txt")
	xorVal, output, _ := set13.XORChar(input, training)
	fmt.Printf("%#U : %s", xorVal, output)
}
