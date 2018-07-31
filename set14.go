package main

import (
	"fmt"
	"os"

	"github.com/shrayolacrayon/cryptopals/set13"
)

func findEncryptedPhrase(filepath string) {
	var max int
	var maxChar rune
	var maxOutput []byte
	var maxInput string
	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	line := make([]byte, 61)
	for {
		_, err := f.Read(line)
		if err != nil {
			break
		}
		input := string(line[:60])
		char, output, sum := set13.XORChar(input)
		if max < sum {
			max = sum
			maxChar = char
			maxOutput = output
			maxInput = input
		}
	}
	fmt.Printf("Input: %s, CHAR: %#U, Phrase: %s", maxInput, maxChar, maxOutput)
	// Input: 7b5a4215415d544115415d5015455447414c155c46155f4058455c5b523f, CHAR: U+0035 '5', Phrase: Now that the party is jumping

}
func run13() {
	// set13
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	xorVal, output, _ := set13.XORChar(input)
	fmt.Printf("%#U : %s", xorVal, output)
}
func main() {
	findEncryptedPhrase("testdata/s4.txt")
}
