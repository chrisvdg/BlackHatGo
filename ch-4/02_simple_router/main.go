package main

import (
	"fmt"
	"net/http"
	"flag"
)

func main() {
	addrPtr := flag.String("l", ":8080", "Listen address")
	flag.Parse()

	r := &router{}
	fmt.Printf("Listening on %s\n", *addrPtr)
	http.ListenAndServe(*addrPtr, r)
}

type router struct {}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/a":
		fmt.Fprint(w, "Executing /a")
	case "/b":
		fmt.Fprint(w, "Executing /b")
	case "/c":
		fmt.Fprint(w, "Executing /c")
	default:
		http.Error(w, fmt.Sprintf("404: %s Not Found", req.URL.Path), 404)
	}
}
