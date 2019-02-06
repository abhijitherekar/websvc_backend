package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

type config struct {
	addr    string
	tlscert string
	tlskey  string
}

func main() {
	var c config
	fmt.Println("Starting the web-app")
	flag.StringVar(&c.addr, "addr", ":4080", "addr on which to listen on")
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/create", create)
	mux.HandleFunc("/snippet", show)

	//create a file server
	fsHandler := http.FileServer(http.Dir("./ui/static"))

	mux.Handle("/static/", http.StripPrefix("/static", fsHandler))

	err := http.ListenAndServe(c.addr, mux)
	log.Fatal(err)
}
