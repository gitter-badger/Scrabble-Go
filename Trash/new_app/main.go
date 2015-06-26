package main

import (
	"fmt"
	"html/template"
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
	http.HandleFunc("/", log_in)
	// log.Fatal(http.ListenAndServe("0.0.0.0:9999", nil))
	fmt.Println("Working")
}

func log_in(rw http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("html/log_in.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(rw, nil)
}
