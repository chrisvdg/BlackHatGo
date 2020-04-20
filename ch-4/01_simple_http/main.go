package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	addrPtr := flag.String("l", ":8080", "Listen address")
	flag.Parse()
	http.HandleFunc("/hello", hello)
	fmt.Printf("Listening on %s\n", *addrPtr)
	http.ListenAndServe(*addrPtr, nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s\n", r.URL.Query().Get("name"))
}
