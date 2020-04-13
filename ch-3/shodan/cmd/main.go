package main

import (
	"fmt"
	"log"
	"os"

	"github.com/chrisvdg/BlackHatGo/ch-3/shodan/shodan"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage: main <searchterm>")
	}
	apiKey := os.Getenv("SHODAN_API_KEY")
	if apiKey == "" {
		log.Fatalln("No API key was found")
	}
	s := shodan.New(apiKey)
	info, err := s.APIInfo()
	if err != nil {
		log.Panicf("Failed to fetch API Info: %s\n", err)
	}
	fmt.Printf(
		"Query Credits: %d\nScan Credits:  %d\n\n",
		info.QueryCredits,
		info.ScanCredits)

	hostSearch, err := s.HostSearch(os.Args[1])
	if err != nil {
		log.Panicf("Failed to search for host %s: %s\n", os.Args[1], err)
	}

	for _, host := range hostSearch.Matches {
		fmt.Printf("%18s%8d\t%s\n", host.IPString, host.Port, host.Hostnames)
	}
}
