package main

import (
	"net/http"

	conn "github.com/Aesir-Development/yugioh-backend/internal/db" // Importing the DB connection package
	// "github.com/Aesir-Development/yugioh-backend/pkg/card" // Importing the card package
	"github.com/gin-gonic/gin"
	"strconv"
	"net/url"
)

type ErrorMessage struct {
	Message string `json:"message"`
}


func main() {
	conn.Setup() // Setting up the DB structure and tables if they don't exist

	r := gin.Default()

	r.GET("/cards", func(c *gin.Context) {
		name := c.Query("name")
		id := c.Query("id")

		if id != "" {

			intID, err := strconv.Atoi(id)

			if err != nil {
				c.JSON(http.StatusBadRequest, "{\"message\": \"Invalid ID\"}")
				return
			}
	
			card := conn.FetchCardByID(intID)
			c.JSON(http.StatusOK, card)
		} else if name != "" {

			// URL decode the name
			decodedName, err := url.QueryUnescape(name)

			if err != nil {
				c.JSON(http.StatusBadRequest, "{\"message\": \"Invalid name\"}")
				return
			}

			card := conn.FetchCardByName(decodedName)
			c.JSON(http.StatusOK, card)
		} else {
			c.JSON(http.StatusBadRequest, ErrorMessage {Message: "Invalid query parameters"})
		}
	})

	// WARNING - This should be removed after first run, it's only to get all cards from the API and save them to the DB
	r.GET("/cards/getcards", func(c *gin.Context) {
		allCards := conn.GetCards()
		conn.SaveCards(allCards)
		c.JSON(http.StatusOK, ErrorMessage {Message: "Cards saved to DB"})
	})
	
	r.Run(":8080")
}