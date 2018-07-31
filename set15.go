package main

import (
	"fmt"

	"github.com/shrayolacrayon/cryptopals/set15"
)

func main() {
	val := set15.RepeatingKeyXOR(stanza, "ICE")
	fmt.Println(hex.EncodeToString(val))
	fmt.Printf("Matched expected: %v", (hex.EncodeToString(val) == expected))
}
