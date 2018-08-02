package set16

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math"
	"os"

	"github.com/shrayolacrayon/cryptopals/set13"
)

const KEYSIZE = 40

func DecryptBase64(filepath string) []byte {
	// read in the whole file
	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	allBytes, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	allBytes, err = base64.StdEncoding.DecodeString(string(allBytes))
	if err != nil {
		panic(err)
	}
	return allBytes
}

// Write a function to compute the edit distance/Hamming distance between two strings.
// The Hamming distance is just the number of differing bits.
func HammingDistance(x, y []byte) int {
	if len(x) != len(y) {
		panic("not the same length")
	}
	var distance int
	for i, xByte := range x {
		bin := fmt.Sprintf("%b", (xByte ^ y[i]))
		for _, char := range bin {
			switch char {
			case '1':
				distance++
			}
		}
	}
	return distance
}

func findMinDistance(distances []float64, ignoreKeys []int) (int, float64) {
	minDistance := math.MaxFloat64
	keysize := 0
	for i := 2; i < KEYSIZE; i++ {
		var ignore bool
		for _, key := range ignoreKeys {
			if i == key {
				ignore = true
			}
		}
		if distances[i] < minDistance && !ignore {
			keysize = i
			minDistance = distances[i]
		}
	}
	return keysize, minDistance
}

func split(allBytes []byte, keysize int) [][]byte {
	var chunk []byte
	chunks := make([][]byte, 0, len(allBytes)/keysize+1)
	for len(allBytes) >= keysize {
		chunk, allBytes = allBytes[:keysize], allBytes[keysize:]
		chunks = append(chunks, chunk)
	}
	//if len(allBytes) > 0 {
	//	chunks = append(chunks, allBytes[:len(allBytes)])
	//}
	return chunks
}

func transpose(matrix [][]byte) [][]byte {
	transposed := make([][]byte, len(matrix[0]))
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			transposed[col] = append(transposed[col], matrix[row][col])
		}
	}
	return transposed
}

func getNormalizedDistance(allBytes []byte, keysize, numBlocks int) float64 {
	var sum float64
	for i := 0; i < numBlocks; i++ {
		start := allBytes[keysize*i : keysize*(i+1)]
		finish := allBytes[keysize*(i+1) : keysize*(i+2)]
		distance := HammingDistance(start, finish)
		normalized := float64(distance) / float64(keysize)
		sum += normalized
	}
	return sum
}

func breakXOR(allBytes []byte, training map[rune]int, keysize int) []byte {
	allLines := split(allBytes, keysize)
	transposed := transpose(allLines)
	xorKey := []rune{}
	for _, line := range transposed {
		char, _, _ := set13.XORChar(hex.EncodeToString(line), training)
		xorKey = append(xorKey, char)
	}
	return []byte(string(xorKey))
}
func BreakXOR(trainingFilepath, filepath string) []byte {
	distances := make([]float64, KEYSIZE)
	allBytes := DecryptBase64(filepath)
	training := set13.CreateTrainingMap(trainingFilepath)

	/*hexString := hex.EncodeToString(allBytes)
	allBytes = []byte(hexString)*/
	for keysize := 3; keysize < KEYSIZE; keysize++ {
		//For each KEYSIZE, take the first KEYSIZE worth of bytes, and the second KEYSIZE worth of bytes,
		//and find the edit distance between them. Normalize this result by dividing by KEYSIZE.
		normalized := getNormalizedDistance(allBytes, keysize, 4)
		distances[keysize] = normalized
	}
	ignoreList := []int{}
	xorKeys := [][]byte{}
	for i := 0; i < 3; i++ {
		keysize, _ := findMinDistance(distances, ignoreList)
		ignoreList = append(ignoreList, keysize)
		xorKey := breakXOR(allBytes, training, keysize)
		xorKeys = append(xorKeys, xorKey)
	}

	// compare the xor keys to english
	var maxSum int
	var winningKey []byte
	for _, key := range xorKeys {
		sum := set13.CompareOutput(key, training)
		if sum > maxSum {
			maxSum = sum
			winningKey = key
		}
	}
	return winningKey

}
