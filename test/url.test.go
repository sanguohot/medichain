package main

import (
	"fmt"
	"net/url"
	"path"
)

func main() {
	//t,_ := url.PathUnescape("http://baidu.com//abcd/efg")
	u, _ := url.Parse("http://baidu.com/")
	fmt.Println(u)
	t := path.Join(u.Path, "/efg")
	fmt.Println(t)
}