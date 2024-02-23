package main

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"

	conn "github.com/Aesir-Development/yugioh-backend/internal/db" // Importing the DB connection package
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

		t := c.Query("type")
		limit := c.Query("limit")
		if (limit == "") {
			c.JSON(http.StatusBadRequest, ErrorMessage {Message: "Invalid query parameters"})
			return
		}

		limitInt, err := strconv.Atoi(limit)
		if (err != nil) {
			c.JSON(http.StatusBadRequest, ErrorMessage {Message: "Invalid limit"})
			return
		}

		// If the type is set, we only want to fetch cards of that type
		if t != "" {


			cards, err := conn.FetchCardByType("Spell Card", limitInt)

			if (err != nil) {
				c.JSON(http.StatusInternalServerError, ErrorMessage {Message: err.Error()})
				return
			}
	
			println(len(cards))
			
			c.JSON(http.StatusOK, cards)
			return
		}

		name := c.Query("name")

		if (name == "") {
			c.JSON(http.StatusBadRequest, ErrorMessage {Message: "Invalid query parameters"})
			return
		}

		escapedName := url.QueryEscape(name)
		escapedName = strings.ReplaceAll(escapedName, "+", " ") // QueryEscape replaces spaces with +, so we need to replace them back to spaces

		println(escapedName, limitInt)

		cards, err := conn.FetchCardsByName(escapedName, limitInt)

		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorMessage {Message: err.Error()})
			return
		}

		c.JSON(http.StatusOK, cards)
	})

	// WARNING - This should be removed after first run, it's only to get all cards from the API and save them to the DB
	r.GET("/unsafe/getcards", func(c *gin.Context) {
		allCards := conn.GetCards()
		conn.SaveCards(allCards)
		c.JSON(http.StatusOK, ErrorMessage {Message: "Cards saved to DB"})
	})
	

	r.Run(":8080")
}