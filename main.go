package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type card struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	r := gin.Default()
	r.GET("/cards", func(c *gin.Context) {
		cards := []card{
			{ID: 1, Name: "card1"},
			{ID: 2, Name: "card2"},
		}
		c.JSON(http.StatusOK, cards)
	})
	r.Run(":8080")
}

// Path: main.go