package main

import (
	"medichain/file"
	"fmt"
	"log"
)

func main() {
	fileBytes, err := file.DownloadFromBigDataCenter("562f2ea6e41020fd7bf5426bd77cd59c")
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("\ndownload file length %d", len(fileBytes))
}