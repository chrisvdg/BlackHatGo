package main

import (
	"bufio"
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
	for {
		reader := bufio.NewReader(conn)
		s, err := reader.ReadString('\n')
		if err == io.EOF {
			log.Printf("%s disconnected\n", conn.RemoteAddr())
			break
		}
		if err != nil {
			log.Printf("Failed to read from %s: %s", conn.RemoteAddr(), err)
			break
		}
		log.Printf("Received %d bytes from %s: %s", len(s), conn.RemoteAddr(), s)

		writer := bufio.NewWriter(conn)
		if _, err := writer.WriteString(s); err != nil {
			log.Printf("Failed to write to %s: %s", conn.RemoteAddr(), err)
			break
		}
		writer.Flush()
	}
}
