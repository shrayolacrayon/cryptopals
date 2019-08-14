package main

import (
	"fmt"

	"cryptopals/set211"
	"cryptopals/util"
)

func main() {
	input, err := util.ReadFile("testdata/someinput.txt")
	if err != nil {
		panic(err)
	}
	for i := 0; i < 10; i++ {
		encrypted, err := set211.EncryptionOracle([]byte(input))
		if err != nil {
			panic(err)
		}
		fmt.Println(set211.DetectMethod(encrypted))
	}
}
