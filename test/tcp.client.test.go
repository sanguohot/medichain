package main
import (
	"net"
	"log"
	"fmt"
)
func main() {
	raddr, err := net.ResolveTCPAddr("tcp", "10.6.250.52:8000")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.DialTCP("tcp", nil, raddr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	buf := make([]byte, 0xffff)
	go func() {
		for {
			n, err := conn.Read(buf)
			if err != nil {
				log.Fatal(err)
			}
			b := buf[:n]
			fmt.Print(string(b))
		}
	}()
	conn.Write([]byte("hello world!"))
}