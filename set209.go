package main

import (
	"fmt"

	"github.com/shrayolacrayon/cryptopals/pkcs7"
)

func main() {
	output := pkcs7.PKCS7([]byte("YELLOW SUBMARINE"), 20)
	fmt.Println(output)
}
