package main

import (
	"flag"
	"fmt"
	"net"
)

func main() {
	addrPtr := flag.String("addr", "scanme.nmap.org", "TCP address to scan")
	minPortPtr := flag.Int("min", 80, "Ports min of range to scan")
	maxPortPtr := flag.Int("max", 0, "Ports max of range to scan")
	verbosePrt := flag.Bool("v", false, "Verbose output")
	flag.Parse()
	addr := *addrPtr
	minPort := *minPortPtr
	maxPort := *maxPortPtr
	verbose := *verbosePrt

	if maxPort == 0 {
		maxPort = minPort
	}
	for i := minPort; i <= maxPort; i++ {
		fullAddr := fmt.Sprintf("%s:%d", addr, i)
		conn, err := net.Dial("tcp", fullAddr)
		if conn != nil {
			conn.Close()
		}
		if err != nil && verbose {
			fmt.Printf("Failed to open TCP connection on %s: %s\n", fullAddr, err)
		}
		if err == nil {
			fmt.Printf("%s is open\n", fullAddr)
		}
	}
}
