package main

import (
	"fmt"

	"github.com/shrayolacrayon/cryptopals/set210"
	"github.com/shrayolacrayon/cryptopals/util"
)

var secret = "YELLOW SUBMARINE"

func main() {
	allBytes, err := util.ReadFile("testdata/challenge210.txt")
	if err != nil {
		panic(err)
	}
	//bytes, err := base64.StdEncoding.DecodeString(string(allBytes))
	//if err != nil {
	//		panic(err)
	//	}

	output, err := set210.CBCOtherDecrypt(allBytes, []byte(secret))
	if err != nil {
		panic(err)
	}
	fmt.Println(string(output))
}
