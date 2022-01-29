package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/heroes", getHeroes)
	router.Run("localhost:8080")
	//fmt.Println(getHeroes()["Hero1"])
}

type Hero struct {
	Name       string         `json:"name"`
	Class      string         `json:"class"`
	Race       string         `json:"race"`
	Gender     string         `json:"gender"`
	Weight     float64        `json:"weight"`
	Height     float64        `json:"height"`
	Hp         int            `json:"hp"`
	Money      float64        `json:"monye"`
	Attributes map[string]int `json:"attributes"`
	Talents    map[string]int `json:"talents"`
}

func getHeroes(c *gin.Context) {
	attr := map[string]int{"attr": 1, "attr2": 2}
	talents := map[string]int{"talent": 1}
	heroes := map[string]Hero{
		"Hero1": Hero{"test", "test", "test", "test", 10, 20, 30, 40, attr, talents}}
	c.IndentedJSON(http.StatusOK, heroes)
}