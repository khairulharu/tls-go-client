package main

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"os"
)

func main() {
	log.SetFlags(log.Lshortfile)

	cert, err := os.ReadFile("/home/khairulaswad/Documents/golang/tls-go/boringssl/cert.pem")
	if err != nil {
		log.Fatalf("Couldn't load file get %s", err)
	}
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(cert)

	conf := &tls.Config{
		RootCAs: certPool,
	}

	conn, err := tls.Dial("tcp", "localhost:5555", conf)
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
}
