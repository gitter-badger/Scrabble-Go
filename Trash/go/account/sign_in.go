package account

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type session struct{}

var Ses session

func (ses session) Create(rw http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("html/log_in.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(rw, nil)
}

func (ses session) Landing_page(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("Landing_page")
	t, err := template.ParseFiles("html/scrabble.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(rw, nil)
}
