package game

import (
	"fmt"
	"github.com/gorilla/mux"
	"honnef.co/go/js/dom"
	"net/http"
)

type verifygame struct{}

var Veri verifygame

func (ver verifygame) Verify(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	player1 := vars["player1"]
	player2 := vars["player2"]

	var alphabets = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "L", "M", "N", "O", "P", "R", "S", "T", "U", "V", "X", "Z"}
	var scores = []int{1, 9, 1, 2, 1, 8, 9, 10, 1, 10, 1, 4, 1, 1, 2, 1, 1, 1, 1, 8, 10, 10}
	var i int
	var j int
	ok := 1
	partial_score := 0
	multiplicator := 1
	fmt.Println(player1, player2, i, j, ok, alphabets, scores, partial_score, multiplicator)
	for i = 1; i < 16; i++ {
		// actual_word := ""
		for j = 1; j < 16; j++ {
			document := dom.GetWindow().Document()
			login := document.GetElementByID("login").(*dom.HTMLDivElement)
			value := login.GetAttribute("value")
			fmt.Println(value)
		}
	}
	http.Redirect(rw, req, "localhost:9999/landing_page", http.StatusFound)
}
