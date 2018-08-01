package set14

import (
	"fmt"
	"os"

	"github.com/shrayolacrayon/cryptopals/set13"
)

func FindEncryptedPhrase(inputFilepath, trainingFilepath string) {
	training := set13.CreateTrainingMap(trainingFilepath)
	var max int
	var maxChar rune
	var maxOutput []byte
	var maxInput string
	f, err := os.Open(inputFilepath)
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
		char, output, sum := set13.XORChar(input, training)
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
