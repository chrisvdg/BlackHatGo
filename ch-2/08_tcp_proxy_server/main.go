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
	portPtr := flag.Int("p", 8001, "Listen port")
	destPtr := flag.String("d", "localhost:8000", "Destination address to proxy to")
	flag.Parse()
	addr := *addrPtr
	port := *portPtr
	destAddr := *destPtr

	listAddr := fmt.Sprintf("%s:%d", addr, port)
	listener, err := net.Listen("tcp", listAddr)
	if err != nil {
		log.Fatalf("Failed to bind to address %s: %s", listAddr, err)
	}
	log.Printf("Listening on %s\n", listAddr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %s", err)
		}

		go handle(conn, destAddr)
	}
}

func handle(src net.Conn, dstAddr string) {
	log.Printf("Received connection from %s\n", src.RemoteAddr())
	defer func() {
		log.Printf("Closing connection with %s\n", src.RemoteAddr())
		src.Close()
	}()
	dst, err := net.Dial("tcp", dstAddr)
	if err != nil {
		log.Printf("Failed to connect to destination %s: %s\n", dstAddr, err)
		return
	}
	defer dst.Close()

	// prevent copy from blocking
	go func() {
		if _, err := io.Copy(dst, src); err != nil {
			log.Printf("Failed to write source to destination: %s\n", err)
		}
	}()

	if _, err := io.Copy(src, dst); err != nil {
		log.Printf("Failed to write destination to source: %s\n", err)
	}
}
