package main

import (
	"flag"
	"fmt"
	"net"
	"sort"
)

func main() {
	addrPtr := flag.String("addr", "scanme.nmap.org", "TCP address to scan")
	minPortPtr := flag.Int("min", 80, "Ports min of range to scan")
	maxPortPtr := flag.Int("max", 0, "Ports max of range to scan")
	verbosePrt := flag.Bool("v", false, "Verbose output")
	workersPtr := flag.Int("workers", 10, "Worker pool size")
	flag.Parse()
	addr := *addrPtr
	minPort := *minPortPtr
	maxPort := *maxPortPtr
	verbose := *verbosePrt
	workers := *workersPtr

	if maxPort == 0 {
		maxPort = minPort
	}
	var openPorts []int
	portsChan := make(chan int, workers)
	resultChan := make(chan int)

	for i := 0; i < cap(portsChan); i++ {
		go worker(addr, portsChan, resultChan, verbose)
	}

	go func() {
		for i := minPort; i <= maxPort; i++ {
			if verbose {
				fmt.Printf("port scan %d scheduled\n", i)
			}
			portsChan <- i
		}
	}()

	for i := minPort; i <= maxPort; i++ {
		port := <-resultChan
		if port != 0 {
			openPorts = append(openPorts, port)
		}
	}
	close(portsChan)
	close(resultChan)

	sort.Ints(openPorts)
	for _, port := range openPorts {
		fullAddr := fmt.Sprintf("%s:%d", addr, port)
		fmt.Printf("%s is open\n", fullAddr)
	}
}

func worker(addr string, ports, result chan int, verbose bool) {
	for port := range ports {
		fullAddr := fmt.Sprintf("%s:%d", addr, port)
		conn, err := net.Dial("tcp", fullAddr)
		if conn != nil {
			conn.Close()
		}
		if err != nil {
			result <- 0
			if verbose {
				fmt.Printf("Failed to open TCP connection on %s: %s\n", fullAddr, err)
			}
		} else {
			result <- port
		}
	}
}
