package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"github.com/abhijitherekar/websvc_backend/pkg/config"
)

type config struct {
	addr    string
	tlscert string
	tlskey  string
}

func main() {
	var c config

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &config.Application{
		Errlog:  errorLog,
		Infolog: infoLog,
	}
	infoLog.Println("Starting the web-app")
	flag.StringVar(&c.addr, "addr", ":4080", "addr on which to listen on")
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/create", app.create)
	mux.HandleFunc("/snippet", app.show)

	//create a file server
	fsHandler := http.FileServer(http.Dir("./ui/static"))

	mux.Handle("/static/", http.StripPrefix("/static", fsHandler))

	//Now instead of piggybaging on the http server, we need our own
	//http server, because, that would help in customizing and adding our
	//own errorlogs.
	svr := http.Server{
		Addr:      c.addr,
		Handler:   mux,
		TLSConfig: nil,
		ErrorLog:  errorLog,
	}
	err := svr.ListenAndServe()
	errorLog.Fatalln(err)
}
