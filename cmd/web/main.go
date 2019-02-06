package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type svrConfig struct {
	addr    string
	tlscert string
	tlskey  string
}

func main() {
	var c svrConfig

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		Errlog:  errorLog,
		Infolog: infoLog,
	}
	infoLog.Println("Starting the web-app")
	flag.StringVar(&c.addr, "addr", ":4080", "addr on which to listen on")
	flag.Parse()

	//Now instead of piggybaging on the http server, we need our own
	//http server, because, that would help in customizing and adding our
	//own errorlogs.
	svr := http.Server{
		Addr:      c.addr,
		Handler:   app.routes(),
		TLSConfig: nil,
		ErrorLog:  errorLog,
	}
	err := svr.ListenAndServe()
	errorLog.Fatalln(err)
}
