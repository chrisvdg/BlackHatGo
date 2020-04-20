package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	addrPtr := flag.String("l", ":8080", "Listen address")
	flag.Parse()

	f := http.HandlerFunc(hello)
	l := logger{Inner: f}

	fmt.Printf("Listening on %s\n", *addrPtr)
	http.ListenAndServe(*addrPtr, &l)
}

type logger struct {
	Inner http.Handler
}

func (l *logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("start")
	l.Inner.ServeHTTP(w, r)
	log.Println("finish")
}

func hello(w http.ResponseWriter, r *http.Request) {
	time.Sleep(1 * time.Second)
	fmt.Fprintf(w, fmt.Sprintf("Hello %s\n", r.URL.Path))
}
