package main

import (
	"fmt"

	"github.com/shrayolacrayon/cryptopals/set210"
)

var secret = "YELLOW SUBMARINE"

func main() {
	output, err := set210.CBC("testdata/challenge210.txt", []byte(secret))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", output)
}
