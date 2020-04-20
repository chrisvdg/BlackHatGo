package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func main() {
	addrPtr := flag.String("l", ":8080", "Listen address")
	flag.Parse()

	r := mux.NewRouter()
	r.HandleFunc("/hello", hello).Methods("GET")

	n := negroni.Classic()
	n.Use(&badAuthorizer{
		Username: "admin",
		Password: "password",
	})
	n.UseHandler(r)

	fmt.Printf("Listening on %s\n", *addrPtr)
	http.ListenAndServe(*addrPtr, n)
}

type contextKey int

const (
	keyUsername contextKey = iota
)

type badAuthorizer struct {
	Username string
	Password string
}

func (b *badAuthorizer) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")
	if username != b.Username && password != b.Password {
		http.Error(w, "Unauthorized", 401)
		return
	}
	ctx := context.WithValue(r.Context(), keyUsername, username)
	r = r.WithContext(ctx)
	next(w, r)
}

func hello(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value(keyUsername).(string)
	fmt.Fprintf(w, "Hi %s\n", username)
}
