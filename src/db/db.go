package db

import (
	"database/sql"
	"encoding/json"
	"log"
)

func GetDBConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/dsa_db")
	if err != nil {
		log.Fatal(err.Error())
	}

	return db, err
}

func GetHero(hero string) data.Hero {
	var (
		name string
		maxHp int
		hp int
		race string
		gender string
		weight float64
		height float64
		money float64
	)

	db, err := GetDBConnection()
	rows, err := db.Query("SELECT * FROM heroes WHERE name = ?", hero)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	attr := map[string]int{}
	t := map[string]int{}

	for rows.Next() {
		err := rows.Scan(&name, &maxHp, &hp, &race, &gender, &weight, &height, &money)
		if err != nil {
			log.Fatal(err)
		}

		json.Unmarshal([]byte(attributes), &attr)
		json.Unmarshal([]byte(talents), &t)
	}
	h := data.Hero
	h.Name = name
	h.MaxHp = maxHp
	h.Hp = hp
	h.Race = race
	h.Gender = gender
	h.Weight = weight
	h.Height = height
	h.Money = money
}

func GetAttribute(name string, attr string) {

}

func GetTalent(name string, tal string) {
	
}

func getAttributes(name string) {

}

func getTalents(name string) {

}

