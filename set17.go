package main

import (
	"fmt"

	"github.com/shrayolacrayon/cryptopals/set17"
)

func main() {
	b := set17.Decrypt("testdata/challenge7.txt")
	fmt.Println(string(b))
}
