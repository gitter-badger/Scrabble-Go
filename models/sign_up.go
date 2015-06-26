package models

type Register struct {
	Id      int
	Player1 string
	Player2 string
	Score   int
}

type Message struct {
	Success string
	Message string
}

type Error struct {
	Error   string
	Message string
}
