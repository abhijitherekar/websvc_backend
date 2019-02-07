package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/abhijitherekar/websvc_backend/pkg/models/mysql"
)

type application struct {
	Errlog   *log.Logger
	Infolog  *log.Logger
	Snippets *mysql.SnippetModel
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
	title := "ABHI"
	content := "db example from golang"
	exp := "7"
	id, err := app.Snippets.Insert(title, content, exp)
	if err != nil {
		app.Errlog.Println("Error inserting in DB", err)
	}
	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
}

func (app *application) show(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		app.Errlog.Println(err.Error())
		http.Error(w, "wrong query", 405)
		return
	}
	sm, err := app.Snippets.Get(id)
	if err != nil {
		app.Errlog.Println(err.Error())
		app.notFound(w)
	}
	fmt.Fprintf(w, "%v", sm)
}
