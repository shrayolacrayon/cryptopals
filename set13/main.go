package main

import (
	"encoding/hex"
	"fmt"
	"os"
)

func createTrainingMap(filePath string) map[rune]int {
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
			fmt.Println(err)
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

func xorChar(input string, trainingMap map[rune]int) (rune, []byte) {
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
	return maxChar, outputMap[maxChar]

}

func main() {
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	trainingMap := createTrainingMap("testdata/11-0.txt")
	xorVal, output := xorChar(input, trainingMap)
	fmt.Printf("%#U : %s", xorVal, output)

}
