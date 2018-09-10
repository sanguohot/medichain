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
	//fmt.Println("now set file add log list...")
	//err = service.SetFileAddLogList()
	//if err != nil {
	//	log.Fatal(err)
	//}
	fmt.Println("now get file add log list...")
	err, fl := service.GetFileAddLogList("", "", "")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fl)
}