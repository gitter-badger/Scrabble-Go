package main

import (
	"github.com/Scrabble-Go/account"
	"github.com/Scrabble-Go/game"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type justFilesFilesystem struct {
	fs http.FileSystem
}

func (fs justFilesFilesystem) Open(name string) (http.File, error) {
	f, err := fs.fs.Open(name)
	if err != nil {
		return nil, err
	}
	return neuteredReaddirFile{f}, nil
}

type neuteredReaddirFile struct {
	http.File
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", account.Ses.Create)
	r.HandleFunc("/landing_page", account.Ses.Landing_page)
	r.HandleFunc("/verify", game.Veri.Verify)
	fs := justFilesFilesystem{http.Dir("./assets/")}
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(fs)))
	http.Handle("/", r)

	// HTTP Listening Port
	log.Println("main : Started : Listening on: http://localhost:9999 ...")
	log.Fatal(http.ListenAndServe("0.0.0.0:9999", nil))
}
