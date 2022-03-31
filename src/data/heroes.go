package data

import (
	"dicebot/db"
	"encoding/json"
	"io"
	"log"
	"github.com/antonholmquist/jason"
)

type Hero struct {
	Name       string         `json:"name"`
	MaxHp      int            `json:"maxHp"`
	Hp         int            `json:"hp"`
	Race       string         `json:"race"`
	Gender     string         `json:"gender"`
	Class      string         `json:"class"`
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
	heroList = nil
	db, err := db.GetDBConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM heroes")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var h Hero

		err = rows.Scan(&h.Name, &h.MaxHp, &h.Hp, &h.Race, &h.Gender, &h.Class, &h.Weight, &h.Height, &h.Money)
		if err != nil {
			log.Fatal(err)
		}

		attJSON, _ := jason.NewObjectFromBytes([]byte(att))
		h.Attributes = attJSON
		talJSON, _ := jason.NewObjectFromBytes([]byte(tal))
		h.Talents = talJSON

		heroList = append(heroList, &h)
	}

	return heroList
}

func GetHeroByName(name string) *Hero {
	db, err := db.GetDBConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var h Hero
	var att string
	var tal string
	err = db.QueryRow("SELECT * FROM heroes WHERE name=?", name).Scan(&h.Name, &h.MaxHp, &h.Hp, &h.Race, &h.Gender, &h.Weight, &h.Height, &h.Money, &att, &tal)
	if err != nil {
		log.Fatal(err)
	}
	attJSON, _ := jason.NewObjectFromBytes([]byte(att))
	h.Attributes = attJSON
	talJSON, _ := jason.NewObjectFromBytes([]byte(tal))
	h.Talents = talJSON

	heroList = append(heroList, &h)

	return &h
}

func GetHeroField(name string, field string) interface{} {
	// if attributes/talents contains field -> get attribute/talent
	
	db, err := db.GetDBConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var f string
	err = db.QueryRow("SELECT ? FROM heroes WHERE name=?", field, name).Scan(&f)
	if err != nil {
		log.Fatal(err)
	}

	return f
}

func GetHeroAttribute(name string, attribute string) jason.Object {
	db, err := db.GetDBConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var a string
	var v string
	err = db.QueryRow("SELECT attribute, value FROM hero_attributes WHERE name=? AND attribute = ?", name, attribute).Scan(&f, &v)
	if err != nil {
		log.Fatal(err)
	}

	// make json response
}

func SetHeroField(name string, field string, value interface{}) {
	db, err := db.GetDBConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.QueryRow("UPDATE heroes SET ? = ? WHERE name = ?", field, value, name)
}

var heroList []*Hero
attributes[8] := [
	"Mut",
	"Klugheit",
	"Intuition",
	"Geschicklichkeit",
	"Charisma",
	"Fingerfertigkeit",
	"Konstitution",
	"Körperkraft"
]
