package main

import (
	"fmt"

	"github.com/shrayolacrayon/cryptopals/set18"
)

func main() {
	encrypted := set18.DecryptFiles("testdata/challenge8.txt")
	fmt.Printf("%x \n", encrypted)
}
