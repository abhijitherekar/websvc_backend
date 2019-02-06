package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type application struct {
	Errlog  *log.Logger
	Infolog *log.Logger
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}
	// Initialize a slice containing the paths to the two files. Note that the // home.page.tmpl file must be the *first* file in the slice.
	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.Errlog.Println(err.Error())
		app.serverError(w, err)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		app.Errlog.Println(err.Error())
		app.serverError(w, err)
		return
	}
	fmt.Fprintf(w, "welcome home")
}

func (app *application) create(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		http.Error(w, "method not supported", 405)
		return
	}
	fmt.Fprintf(w, "create home")
}

func (app *application) show(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Header().Set("Allow", "GET")
		http.Error(w, "method not supported", 405)
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		app.Errlog.Println(err.Error())
		http.Error(w, "wrong query", 405)
		return
	}

	fmt.Fprintf(w, "show home %d", id)
}
