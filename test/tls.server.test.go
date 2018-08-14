package main
import (
	"bufio"
	"crypto/tls"
	"log"
	"net"
)
func main() {
	cert, err := tls.X509KeyPair([]byte(`
-----BEGIN CERTIFICATE-----
MIICrTCCAZWgAwIBAgIJAId9WK/1W0QMMA0GCSqGSIb3DQEBCwUAMGYxCzAJBgNV
BAMMAkNOMQswCQYDVQQGEwJDTjESMBAGA1UECAwJR3Vhbmdkb25nMREwDwYDVQQH
DAhTaGVuemhlbjETMBEGA1UECgwKRklTQ08tQkNPUzEOMAwGA1UECwwFYWdlbnQw
HhcNMTgwODAyMDkzMzQ0WhcNMTgwOTAxMDkzMzQ0WjB4MS0wKwYDVQQDDCRDTi9l
bWFpbEFkZHJlc3M9c2VydmljZUBmaXNjby5jb20uY24xCzAJBgNVBAYTAkNOMRIw
EAYDVQQIDAlHdWFuZ2RvbmcxETAPBgNVBAcMCFNoZW56aGVuMRMwEQYDVQQKDApG
SVNDTy1CQ09TMFYwEAYHKoZIzj0CAQYFK4EEAAoDQgAEKkDj10+mB1Zs7/iK3SEi
X4wKj3vhra8KDDscEtvwb0sgily0JIAQiu/ZqBb2hYvNQSpvK9TDFqVe7kEMHbHY
I6MaMBgwCQYDVR0TBAIwADALBgNVHQ8EBAMCBeAwDQYJKoZIhvcNAQELBQADggEB
AC5oCYmU/LtZ7DszzDlQjmoDIegiC9v5u7RFP9eWRaVxS8DX2XOBWwYemFgsmr6q
7D/6qsUpw35pRRnXp8TEGo6i5GFOED8VmuZPbmfGRfUg4u4XmJPYuRa9IAQqGnIT
p1t49CFr3GnIvGpvmJrMXg8ZHfPSo7Bbnz5PUOXwjWSQHfrbc/S0TsarA9WIWeEE
7mbJOaGwFhzNzJtUW3k0doc2L3KSZfMm21uDvnZHqeRbDy2KMEl60p56wY9V2tgo
J4jj+P+EDUqcLfx1aO+GjLf5fDf4aQBZhxWPJQirY/tB84R8u6wN7j/DXgbYcgZS
psYHgZLCq8yb38ruFRYqw08=
-----END CERTIFICATE-----
`), []byte(`
-----BEGIN PRIVATE KEY-----
MIGEAgEAMBAGByqGSM49AgEGBSuBBAAKBG0wawIBAQQgiwGaJ0Wpf6tHtINzoOL3
kN9SydBc6HFPKewV4Cw+nQ+hRANCAAQqQOPXT6YHVmzv+IrdISJfjAqPe+GtrwoM
OxwS2/BvSyCKXLQkgBCK79moFvaFi81BKm8r1MMWpV7uQQwdsdgj
-----END PRIVATE KEY-----
`))
	if err != nil {
		log.Println(err)
		return
	}
	config := &tls.Config{Certificates: []tls.Certificate{cert}}
	ln, err := tls.Listen("tcp", ":8822", config)
	if err != nil {
		log.Println(err)
		return
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)
	}
}
func handleConn(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)
	for {
		msg, err := r.ReadString('\n')
		if err != nil {
			log.Println(err)
			return
		}
		println(msg)
		n, err := conn.Write([]byte("world\n"))
		if err != nil {
			log.Println(n, err)
			return
		}
	}
}