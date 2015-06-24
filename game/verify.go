package game

import (
	"fmt"
	// "github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/dom"
	"net/http"
)

type verifygame struct{}

var Veri verifygame

func (ver verifygame) Verify(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("I am in Verify")
	var alphabets = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "L", "M", "N", "O", "P", "R", "S", "T", "U", "V", "X", "Z"}
	var scores = []int{1, 9, 1, 2, 1, 8, 9, 10, 1, 10, 1, 4, 1, 1, 2, 1, 1, 1, 1, 8, 10, 10}
	fmt.Println(alphabets, scores)
	// var i int
	// var j int
	// ok := 1
	d := dom.GetWindow().Document()
	score_player1 := d.GetElementById("score_player1")
	score_player2 := d.GetElementById("score_player2")
	fmt.Println(score_player1, score_player2)
}
