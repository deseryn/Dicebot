package routes

import (
	"dicebot/data"
	"log"
	"net/http"
)

type Heroes struct {
	l *log.Logger
}

func NewHeroes(l *log.Logger) *Heroes {
	return &Heroes{l}
}

func (hs *Heroes) GetAllHeroes(rw http.ResponseWriter, r *http.Request) {
	hs.l.Println("Handle GET Heroes")

	// fetch all heroes fomr the db
	heroList := data.GetHeroes()

	// serialize the list to JSON
	err := heroList.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
