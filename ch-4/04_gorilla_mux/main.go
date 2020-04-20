package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	addrPtr := flag.String("l", ":8080", "Listen address")
	flag.Parse()

	r := mux.NewRouter()

	r.HandleFunc("/foo", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "Hello Foo")
	}).Methods("GET").Host("localhost") // Route only matches when calling the server using localhost, 127.0.0.1 will return 404

	r.HandleFunc("/users/{user}", func(w http.ResponseWriter, req *http.Request) {
		user := mux.Vars(req)["user"]
		fmt.Fprintf(w, "Hello %s\n", user)
	}).Methods("Get")

	// Only match path when user is lowercase
	r.HandleFunc("/lowerusers/{user:[a-z]+}", func(w http.ResponseWriter, req *http.Request) {
		user := mux.Vars(req)["user"]
		fmt.Fprintf(w, "Hello %s\n", user)
	}).Methods("Get")

	fmt.Printf("Listening on %s\n", *addrPtr)
	http.ListenAndServe(*addrPtr, r)
}
