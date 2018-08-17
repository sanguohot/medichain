package main
import (
	"github.com/spacemonkeygo/openssl"
	"flag"
	"os"
	"log"
	"fmt"
	"encoding/json"
)
const bufferSize = 0xffff
type jsonrpcMessage struct {
	Version string          `json:"jsonrpc"`
	ID      uint64 `json:"id,omitempty"`
	Method  string          `json:"method,omitempty"`
	Params  json.RawMessage `json:"params,omitempty"`
	//Result  json.RawMessage `json:"result,omitempty"`
}
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
		log.Fatal("ccccc ===> ", err)
	}
	err = ctx.LoadVerifyLocations(*caFile, "")
	if err != nil {
		log.Fatal("bbbbb ===> ", err)
	}
	conn, err := openssl.Dial("tcp", "10.6.250.58:8822", ctx, openssl.InsecureSkipHostVerification)
	if err != nil {
		log.Fatal("aaaaa ===> ", err)
	}
	buf := make([]byte, bufferSize)
	go func() {
		for {
			n, err := conn.Read(buf)
			if err != nil {
				log.Fatal("eeeee ===> ", err)
			}

			b := buf[:n]
			fmt.Print(string(b))
		}
	}()

	//var params = [1] string {"newHeads"}
	//paramsBytes, err := json.Marshal(params)
	//if err != nil {
	//	log.Fatal("111111 ===> ", err)
	//}
	msg := &jsonrpcMessage{
		ID: 1,
		Version: "2.0",
		Method: "eth_subscribe",
		//Params: paramsBytes,
	}
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		log.Fatal("222222 ===> ", err)
	}
	n, err := conn.Write(msgBytes)
	if err != nil {
		log.Fatal("333333 ===> ", err)
	}
	fmt.Printf("write bytes %d\n", n)
	fmt.Println(msg)
	os.Stdin.Read(make([]byte,1))
}