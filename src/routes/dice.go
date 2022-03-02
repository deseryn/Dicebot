package routes

import (
	"dicebot/data"

	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// TODO add api for dice rolls

type Dice struct {
	l *log.Logger
}

func NewDiceRoute(l *log.Logger) *Dice {
	return &Dice{l}
}

func (dr *Dice) RollDice(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	face, err := strconv.Atoi(vars["face"])
	if err != nil {
		http.Error(rw, "Unable to convert 'face'", http.StatusBadRequest)
		return
	}

	count, err := strconv.Atoi(vars["count"])
	if err != nil {
		http.Error(rw, "Unable to convert 'count'", http.StatusBadRequest)
		return
	}

	diceRolls := data.Roll(face, count)

	dr.l.Println("Roll d", face, count, "times: ", diceRolls)

	err = diceRolls.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
