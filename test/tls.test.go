package main

import (
	"fmt"
	"crypto/tls"
)

func main() {
	certFile := "e:/evan/goland/src/medichain/key/sdk.crt"
	keyFile := "e:/evan/goland/src/medichain/key/sdk.key"
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cert)
}