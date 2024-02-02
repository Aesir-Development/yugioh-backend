package main

import (
	"net/http"
	"net/url"
	"strconv"

	conn "github.com/Aesir-Development/yugioh-backend/internal/db" // Importing the DB connection package
	"github.com/Aesir-Development/yugioh-backend/pkg/card"
	"github.com/gin-gonic/gin"
)

type ErrorMessage struct {
	Message string `json:"message"`
}


func main() {
	err := conn.Setup() // Setting up the DB structure and tables if they don't exists
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	r.GET("/cards", func(c *gin.Context) {
		name := c.Query("name")
		limit := c.Query("limit")

		if (name == "" || limit == "") {
			c.JSON(http.StatusBadRequest, ErrorMessage {Message: "Invalid query parameters"})
			return
		}

		limitInt, err := strconv.Atoi(limit)
		if err != nil {
			c.JSON(http.StatusBadRequest, ErrorMessage {Message: "Invalid limit"})
			return
		}

		escapedName := url.QueryEscape(name)

		cards := GetCards(escapedName, limitInt)

		c.JSON(http.StatusOK, cards)

		// TODO - Fetch cards by name and limit
	})

	// WARNING - This should be removed after first run, it's only to get all cards from the API and save them to the DB
	r.GET("/cards/getcards", func(c *gin.Context) {
		allCards := conn.GetCards()
		conn.SaveCards(allCards)
		c.JSON(http.StatusOK, ErrorMessage {Message: "Cards saved to DB"})
	})
	

	r.Run(":8080")
}

func GetCards(name string, limit int) []card.Card {
	cards := conn.FetchCardsByName(name, limit)
	return cards
}