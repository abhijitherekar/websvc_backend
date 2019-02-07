package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/abhijitherekar/websvc_backend/pkg/models/mysql"

	//the underscroe signifies that only the init func of the package is inhireted
	_ "github.com/go-sql-driver/mysql"
)

type svrConfig struct {
	addr    string
	tlscert string
	tlskey  string
	dbConf  string
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
	flag.StringVar(&c.dbConf, "dsn", "root:@/snippetbox?parseTime=true", "mySql conf")
	flag.Parse()

	db, err := openDB(c.dbConf)
	if err != nil {
		errorLog.Fatalln("Cannot open the DB", err)
	}
	defer db.Close()
	app.Snippets = &mysql.SnippetModel{Db: db}
	//Now instead of piggybaging on the http server, we need our own
	//http server, because, that would help in customizing and adding our
	//own errorlogs.
	svr := http.Server{
		Addr:      c.addr,
		Handler:   app.routes(),
		TLSConfig: nil,
		ErrorLog:  errorLog,
	}
	err = svr.ListenAndServe()
	errorLog.Fatalln(err)
}

func openDB(dns string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dns)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
