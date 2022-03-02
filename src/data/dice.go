package data

import (
	"encoding/json"
	"io"
	"math/rand"
)

type DiceRolls []int

func Roll(faces int, count int) DiceRolls {
	rolls := []int{}

	for i := 0; i < count; i++ {
		roll := rand.Intn(faces-1) + 1
		rolls = append(rolls, roll)
	}

	return rolls
}

func (dr *DiceRolls) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(dr)
}
