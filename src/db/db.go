package db

import (
	"database/sql"
	"dicebot/data"
	"encoding/json"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func DBConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/dsa_db")
	if err != nil {
		log.Fatal("Error: Could not open database connection", err)
	}
	//defer db.Close()

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
		attributes string
		talents string
	)

	db, err := DBConnection()
	rows, err := db.Query("SELECT * FROM heroes WHERE name = ?", hero)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	attr := map[string]int{}
	t := map[string]int{}

	for rows.Next() {
		err := rows.Scan(&name, &maxHp, &hp, &race, &gender, &weight, &height, &money, &attributes, &talents)
		if err != nil {
			log.Fatal(err)
		}

		json.Unmarshal([]byte(attributes), &attr)
		json.Unmarshal([]byte(talents), &t)
	}
	return data.Hero{
		name, 
		maxHp, 
		hp, 
		race, 
		gender, 
		weight, 
		float64(height), 
		float64(money), 
		attr,
		t,
		}
}