package main

import (
	"fmt"

	"github.com/shrayolacrayon/cryptopals/ecb"
)

func main() {
	b := ecb.Decrypt("testdata/challenge7.txt")
	fmt.Println(string(b))
}
