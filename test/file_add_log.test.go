package main

import (
	"fmt"
	"log"
	"medichain/datacenter"
	"medichain/service"
)


func main() {
	err, max := datacenter.SqliteGetMaxBlockNum()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("datacenter.SqliteGetMaxBlockNum ===>", max)
	err, detail := service.GetFileAddLogDetail("568759d5-5d90-4ec8-b7aa-6743a0004f2f")
	fmt.Println("service.GetFileAddLogDetail ===>", detail)
	fmt.Println("now set file add log list...")
	err = service.SetFileAddLogList()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("now get file add log list...")
	err, fl := service.GetFileAddLogList("", "", "")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fl)
}