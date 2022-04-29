package data

import (
	"database/sql"
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

func (h *Hero) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(h)
}

func GetHeroInfo(name string) Hero {
	db, err := db.GetDBConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	h := GetHero(name, db)
	//h.Attributes  = getHeroAttributes(h, db)
	//h.Talents = getHeroTalents(h)
}

func GetHero(name string, db *sql.DB) Hero {
	var h Hero
	err := db.QueryRow("SELECT * FROM heroes WHERE name= ?", name).Scan(&h.Name, &h.MaxHp, &h.Hp, &h.Race, &h.Gender, &h.Weight, &h.Height, &h.Money)
	if err != nil {
		log.Fatal(err)
	}

	return h
}

func GetHeroField(name string, field string) interface{} {	
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

func GetHeroAttribute(name string, attribute string) int {
	db, err := db.GetDBConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var v int
	err = db.QueryRow("SELECT value FROM hero_attributes WHERE name=? AND attribute = ?", name, attribute).Scan(&v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func GetHeroTalent(name string, tal string) int {
	db, err := db.GetDBConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var v int
	err = db.QueryRow("SELECT value FROM hero_talents WHERE name= ? AND talent = ?", name, tal).Scan(&v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

func GetTalent(tal string) Talent {
	db, err := db.GetDBConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var t Talent
	err = db.QueryRow("SELECT name, attr1, attr2, attr3 FROM talents WHERE name=?", tal).Scan(&t.Name, &t.Attr1, &t.Attr2, &t.Attr3)
	if err != nil {
		log.Fatal(err)
	}

	return t
}

func SetHeroField(name string, field string, value interface{}) {
	db, err := db.GetDBConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.QueryRow("UPDATE heroes SET ? = ? WHERE name = ?", field, value, name)
}

func SetHeroAttribute(name string, field string, value interface{}) {
	db, err := db.GetDBConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.QueryRow("UPDATE hero_attributes SET ? = ? WHERE name = ?", field, value, name)
}

func SetHeroTalent(name string, field string, value interface{}) {
	db, err := db.GetDBConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.QueryRow("UPDATE hero_talents SET ? = ? WHERE name = ?", field, value, name)
}

func getHeroAttributes(h Hero, db *sql.DB) Attributes {
	var a Attributes
	db.QueryRow("SELECT attribute, value FROM hero_attributes WHERE name = ?", h.Name).Scan(&a.Courage, &a.Cleverness, &a.Intuition, &a.Agility, &a.Charisma, &a.Dexterity, &a.Constitution, &a.Strength)
	
	return a
}

func getHeroTalents(h Hero, db *sql.DB) {
	// TODO Talents
	//a := Attributes
	//db.QueryRow("SELECT talent, value FROM hero_attributes WHERE name = ?", h.Name).Scan(&a.Courage, &a.Cleverness, &a.Intuition, &a.Agility, &a.Charisma, &a.Dexterity, &a.Constitution, &a.Strength)
	
	//return a
}

var heroList []*Hero

type Talent struct {
	Name string `json: "name"`
	Attr1 int `json: "attr1"`
	Attr2 int `json: "attr2"`
	Attr3 int `json: "attr3"`
}

type Attributes struct {
	Courage int `json: "mut"`
	Cleverness int `json: "klugheit"`
	Intuition int `json: ""`
	Agility int `json: ""`
	Charisma int `json: ""`
	Dexterity int `json: ""`
	Constitution int `json: ""`
	Strength int `json: ""`
}
