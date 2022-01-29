package main

import "fmt"

func main() {
	fmt.Println("success")
}

type Hero struct{
	Name string `json:"name"`
	Class string `json:"class"`
	Race string `json:"race"`
	Gender string `json:"gender"`
	Weight float64 `json:"weight"`
	Height float64 `json:"height"`
	Hp int `json:"hp"`
	Money float64 `json:"monye"`
	Attributes map[string]int `json:"attributes"`
	Talents map[string]int `json:"talents"`
}