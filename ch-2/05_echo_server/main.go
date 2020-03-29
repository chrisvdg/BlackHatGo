package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	addrPtr := flag.String("a", "", "Listen address")
	portPtr := flag.Int("p", 8000, "Listen port")
	flag.Parse()
	addr := *addrPtr
	port := *portPtr

	listenAddr := fmt.Sprintf("%s:%d", addr, port)
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatalf("Unable to listen on %s: %s\n", addr, err)
	}
	log.Printf("Listening on %s\n", listenAddr)

	for {
		conn, err := listener.Accept()
		log.Printf("Received connection from %s\n", conn.RemoteAddr())
		if err != nil {
			log.Printf("Unable to accept connection from %s : %s", conn.RemoteAddr(), err)
		}
		go echo(conn)
	}
}

func echo(conn net.Conn) {
	defer conn.Close()

	b := make([]byte, 512)
	for {
		size, err := conn.Read(b[0:])
		if err == io.EOF {
			log.Printf("%s disconnected\n", conn.RemoteAddr())
			break
		}
		if err != nil {
			log.Printf("Unexpected error from %s: %s\n", conn.RemoteAddr(), err)
			break
		}
		readLine := strings.TrimSuffix(string(b), "\n")
		fmt.Printf("%+v\n", readLine)
		log.Printf("Received %d bytes from %s: %s", size, conn.RemoteAddr(), readLine)
		if _, err := conn.Write(b[0:size]); err != nil {
			log.Printf("Unable to write to %s: %s\n", conn.RemoteAddr(), err)
			break
		}
	}
}
