package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
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
	var wg sync.WaitGroup
	portsChan := make(chan int, workers)

	for i := 0; i < cap(portsChan); i++ {
		go worker(addr, portsChan, &wg, verbose)
	}

	for i := minPort; i <= maxPort; i++ {
		if verbose {
			fmt.Printf("port scan %d scheduled\n", i)
		}
		wg.Add(1)
		portsChan <- i
	}
	wg.Wait()
	close(portsChan)
}

func worker(addr string, ports chan int, wg *sync.WaitGroup, verbose bool) {
	for port := range ports {
		fullAddr := fmt.Sprintf("%s:%d", addr, port)
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
		wg.Done()
	}
}
