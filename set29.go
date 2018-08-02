package main

import (
	"fmt"

	"github.com/shrayolacrayon/cryptopals/set29"
)

func main() {
	output := set29.PKCS7([]byte("YELLOW SUBMARINE"), 20)
	fmt.Println(output)
}
