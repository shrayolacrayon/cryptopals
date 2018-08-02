package main

import (
	"fmt"

	"github.com/shrayolacrayon/cryptopals/set16"
)

func main() {

	// test hamming distance
	input1 := []byte("this is a test")
	input2 := []byte("wokka wokka!!!")
	expectedDistance := 37
	distance := set16.HammingDistance(input1, input2)
	fmt.Printf("Distance is %d, expected %d - %v \n", distance, expectedDistance, (distance == expectedDistance))

	key, err := set16.BreakXOR("testdata/11-0.txt", "testdata/challenge6.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(key))
}
