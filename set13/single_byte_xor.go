package set13

import (
	"encoding/hex"
	"os"
)

func CreateTrainingMap(filePath string) map[rune]int {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	total := 0
	totalMap := map[rune]int{}
	bytez := make([]byte, 1024)
	for {
		_, err := f.Read(bytez)

		if err != nil {
			break
		}
		for _, b := range string(bytez) {
			totalMap[b]++
			total++
		}
	}
	return totalMap
}

func compare(output []byte, training map[rune]int) int {
	var sum int
	var total int
	for _, b := range string(output) {
		// look up the frequencies and multiply by the total, find the difference
		if v, ok := training[b]; ok {
			sum += v
			total++
		}
	}
	return sum
}

func XORChar(input string, training map[rune]int) (rune, []byte, int) {

	decodedInput, err := hex.DecodeString(input)
	outputMap := map[rune][]byte{}
	if err != nil {
		panic(err)
	}

	var maxChar rune
	var maxSum int

	for i := 0; i < 256; i++ {
		// xor against each character in the decoded
		output := make([]byte, len(decodedInput))
		for j, d := range decodedInput {
			output[j] = byte(i) ^ d
		}
		outputMap[rune(i)] = output
		sum := compare(output, trainingMap)
		if sum > maxSum {
			maxChar = rune(i)
			maxSum = sum
		}
	}
	return maxChar, outputMap[maxChar], maxSum

}
