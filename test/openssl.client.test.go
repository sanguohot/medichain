package main
import (
	"github.com/spacemonkeygo/openssl"
	"flag"
	"os"
	"log"
)

func main() {
	var (
		certFile = flag.String("cert", "/opt/conf/sdk.crt", "A PEM eoncoded certificate file.")
		keyFile  = flag.String("key", "/opt/conf/sdk.key", "A PEM encoded private key file.")
		caFile   = flag.String("CA", "/opt/conf/ca.crt", "A PEM eoncoded CA's certificate file.")
		//certFile = flag.String("cert", "e:/evan/goland/src/medichain/key/sdk.crt", "A PEM eoncoded certificate file.")
		//keyFile  = flag.String("key", "e:/evan/goland/src/medichain/key/sdk.key", "A PEM encoded private key file.")
		//caFile   = flag.String("CA", "e:/evan/goland/src/medichain/key/ca.crt", "A PEM eoncoded CA's certificate file.")
	)
	//ctx, err := openssl.NewCtx()
	//if err != nil {
	//	log.Fatal(err)
	//}
	ctx, err := openssl.NewCtxFromFiles(*certFile, *keyFile)
	if err != nil {
		log.Fatal(err)
	}
	err = ctx.LoadVerifyLocations(*caFile, "")
	if err != nil {
		log.Fatal(err)
	}
	_, err = openssl.Dial("tcp", "10.6.250.54:8822", ctx, openssl.InsecureSkipHostVerification)
	if err != nil {
		log.Fatal(err)
	}
	os.Stdin.Read(make([]byte,1))
}