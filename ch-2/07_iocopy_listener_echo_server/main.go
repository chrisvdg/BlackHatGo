package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
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
	if _, err := io.Copy(conn, conn); err != nil {
		log.Printf("Something went wrong copying io: %s\n", err)
	}
	log.Printf("Closing connection with %s\n", conn.RemoteAddr())
}
