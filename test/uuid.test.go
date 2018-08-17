package main

import (
	"fmt"
	"github.com/google/uuid"
	"log"
)

func main() {
	id := uuid.New()
	fmt.Println(len(id))
	fmt.Println(id)
	fmt.Println(id.String())
	id2, err := uuid.Parse(id.String())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(id2))
	fmt.Println(id2.String())
	for i,byte := range id {
		fmt.Printf("\n%d ===> %x", i, byte)
	}
}