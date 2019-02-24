package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/create", app.create)
	mux.HandleFunc("/snippet", app.show)

	//create a file server
	fsHandler := http.FileServer(http.Dir("./ui/static"))

	mux.Handle("/static/", http.StripPrefix("/static", fsHandler))
	return secureHeader(mux)
}
