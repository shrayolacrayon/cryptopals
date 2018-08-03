package main

import (
	"fmt"

	"github.com/shrayolacrayon/cryptopals/set209"
)

func main() {
	output := set209.PKCS7([]byte("YELLOW SUBMARINE"), 20)
	fmt.Println(output)
}
