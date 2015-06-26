package account

import (
	"fmt"
	"log"
	"net/http"
)

type registrationController struct{}

var register registrationController

func (reg register) Create(rw http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)

	var r models.Register

	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(body, &r)
	if err != nil {
		panic(err)
	}

	db, err := sql.Open("postgres", "password=password host=localhost dbname=scrabble_dev sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	users, err := db.Exec("CREATE TABLE IF NOT EXISTS users (id int, player_name varchar(255), score int DEFAULT 0)")
	if err != nil {
		log.Fatal(err)
	}

}
