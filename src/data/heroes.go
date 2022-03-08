package data

import (
	"encoding/json"
	"io"
)

type Hero struct {
	Name       string         `json:"name"`
	MaxHp      int            `json:"maxHp"`
	Hp         int            `json:"hp"`
	Race       string         `json:"race"`
	Gender     string         `json:"gender"`
	Weight     float64        `json:"weight"`
	Height     float64        `json:"height"`
	Money      float64        `json:"money"`
	Attributes map[string]int `json:"attributes"`
	Talents    map[string]int `json:"talents"`
}

// typedef for list of heroes
type Heroes []*Hero

func (h *Heroes) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(h)
}

func GetHeroes() Heroes {
	// TODO this should come from a db
	return dbMock
}

var dbMock = Heroes{
	&Hero{
		ID:     1,
		Name:   "Jago",
		MaxHp:  40,
		Hp:     40,
		Race:   "human",
		Gender: "male",
		Money:  9999,
	},
	&Hero{
		ID:     2,
		Name:   "Estor",
		MaxHp:  29,
		Hp:     20,
		Race:   "human",
		Gender: "male",
		Money:  10,
	},
}
