package main
import (
	"crypto/tls"
	"log"
	"io/ioutil"
	"crypto/x509"
	"flag"
	"os"
)
var (
	certFile = flag.String("cert", "/opt/conf/sdk.crt", "A PEM eoncoded certificate file.")
	keyFile  = flag.String("key", "/opt/conf/sdk.key", "A PEM encoded private key file.")
	caFile   = flag.String("CA", "/opt/conf/ca.crt", "A PEM eoncoded CA's certificate file.")
)
func main() {
	// Load client cert
	// Load client cert
	cert, err := tls.LoadX509KeyPair(*certFile, *keyFile)
	if err != nil {
		log.Fatal(err)
	}

	// Load CA cert
	caCert, err := ioutil.ReadFile(*caFile)
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCertPool,
		InsecureSkipVerify: true,
		MinVersion: tls.VersionSSL30,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_RC4_128_SHA,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA,
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
		},
	}
	conn, err := tls.Dial("tcp", "localhost:8822", tlsConfig)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	n, err := conn.Write([]byte("hello\n"))
	if err != nil {
		log.Println(n, err)
		return
	}
	buf := make([]byte, 100)
	n, err = conn.Read(buf)
	if err != nil {
		log.Println(n, err)
		return
	}
	println(string(buf[:n]))
	os.Stdin.Read(make([]byte,1))
}