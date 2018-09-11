package main

import (
	"fmt"
	"log"
	"github.com/sanguohot/medichain/datacenter"
	"github.com/sanguohot/medichain/service"
)


func main() {
	err, total := datacenter.SqliteGetFileAddLogTotal("", "", "")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("datacenter.SqliteGetFileAddLogTotal ===>", total)
	err, detail := service.GetFileAddLogDetail("568759d5-5d90-4ec8-b7aa-6743a0004f2f")
	fmt.Println("service.GetFileAddLogDetail ===>", detail)
	fmt.Println("now set file add log list...")
	err = service.SetFileAddLogList()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("now get file add log list...")
	err, fl := service.GetFileAddLogList("", "", "", "")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fl)
}