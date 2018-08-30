package main

import (
	"medichain/service"
	"log"
	"fmt"
)


func main() {
	fmt.Println("now set file add log list...")
	err := service.SetFileAddLogList()
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