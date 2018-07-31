package main

import (
	"fmt"
)

func main() {
	// 写法1
	list := []int{100, 200}
	fmt.Println("old list ===>", list)
	list[0], list[1] = list[1], list[0]
	fmt.Println("new list ===>", list)
	// 写法2
	a := 100
	b := 200
	a = a+b
	b = a-b
	a = a-b
	fmt.Printf("a ===> %d, b ===> %d", a, b)
}