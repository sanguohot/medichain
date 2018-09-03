package main

import (
	"fmt"
	"go-ethereum/common/hexutil"
	"log"
	"net/http"
	"golang.org/x/text/transform"
	"golang.org/x/text/encoding/simplifiedchinese"
	"strings"
	"io/ioutil"
	"medichain/util"
)
func DecodeToGBK(utf8Str string) (dst string, err error) {
	var trans transform.Transformer = simplifiedchinese.GBK.NewEncoder()
	var reader *strings.Reader = strings.NewReader(utf8Str)
	var transReader *transform.Reader = transform.NewReader(reader, trans)
	bytes, err := ioutil.ReadAll(transReader)
	if err != nil {
		return
	}
	dst = string(bytes)
	return
}

func EncodeFromGBK(gbkStr string) (utf8Str string, err error) {
	var trans transform.Transformer = simplifiedchinese.GBK.NewDecoder()
	var reader *strings.Reader = strings.NewReader(gbkStr)
	var transReader *transform.Reader = transform.NewReader(reader, trans)
	bytes, err := ioutil.ReadAll(transReader)
	if err != nil {
		return
	}
	utf8Str = string(bytes)
	return
}
func main() {
	bytes, err := hexutil.Decode("0x68656c6c6f20776f726c64210d0aced2cac7bae3dec8313131313232323232")
	if err != nil {
		log.Fatal(err)
	}
	gbkStr := string(bytes)
	fmt.Println("http.DetectContentType ===>", http.DetectContentType(bytes))
	fmt.Println("gbkStr ===>", gbkStr)
	gbkToUtf8Str, err := EncodeFromGBK(gbkStr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("gbkToUtf8Str ===>", gbkToUtf8Str)
	fmt.Println(util.DetectContentType(bytes))
}