package main

import (
	"github.com/sanguohot/medichain/util"
	"fmt"
)


func main() {
	content := "hello"
	fmt.Println(util.Sha256Hash([]byte(content)).Hex())
	// 0x2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824
	// echo -n "hello" | openssl dgst -sha256
	// (stdin)= 2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824
}