package main

import (
	"fmt"

	"github.com/shrayolacrayon/cryptopals/set18"
)

func main() {
	encrypted, err := set18.DecryptFiles("testdata/challenge8.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%x \n", encrypted)
}
