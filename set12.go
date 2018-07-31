package main

import (
	"log"

	"github.com/shrayolacrayon/cryptopals/set12"
)

func main() {

	input := "1c0111001f010100061a024b53535009181c"
	xorVal := "686974207468652062756c6c277320657965"

	s, err := set12.FixedXOR(input, xorVal)
	if err != nil {
		log.Printf("error: %s", err)
		return
	}
	log.Printf("val: %s", s)
	//expected := "746865206b696420646f6e277420706c6179"

}
