package main

import (
	"encoding/hex"
	"fmt"
	"log"
)

func fixedXOR(input1, input2 string) ([]byte, error) {
	decode1, err := hex.DecodeString(input1)
	if err != nil {
		return nil, err
	}
	decode2, err := hex.DecodeString(input2)
	if err != nil {
		return nil, err
	}
	if len(decode1) != len(decode2) {
		return nil, fmt.Errorf("not the same length strings")
	}

	newBytes := make([]byte, len(decode1))
	for i, i1 := range decode1 {
		newBytes[i] = i1 ^ decode2[i]
	}

	return newBytes, nil
}

func main() {

	input := "1c0111001f010100061a024b53535009181c"
	xorVal := "686974207468652062756c6c277320657965"

	s, err := fixedXOR(input, xorVal)
	if err != nil {
		log.Printf("error: %s", err)
		return
	}
	log.Printf("val: %s", s)
	expected := "746865206b696420646f6e277420706c6179"
}
