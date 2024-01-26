package main

import (
	"net/http"

	conn "github.com/Aesir-Development/yugioh-backend/internal/db" // Importing the DB connection package
	"github.com/Aesir-Development/yugioh-backend/pkg/card" // Importing the card package
	"github.com/gin-gonic/gin"
)


func main() {
	conn.Setup() // Setting up the DB structure and tables if they don't exist

	r := gin.Default()

	// TODO - Add support for fetching specific cards from the DB
	r.GET("/cards", func(c *gin.Context) {
		cards := []card.Card{
			{
				Name: "Blue-Eyes White Dragon",
				Description: "This legendary dragon is a powerful engine of destruction.",
				Attack: 3000,
				Defense: 2500,
				Level: 8,
				Attribute: "LIGHT",
				Type: "Dragon",
			},
			{
				Name: "Dark Magician",
				Description: "The ultimate wizard in terms of attack and defense.",
				Attack: 2500,
				Defense: 2100,
				Level: 7,
				Attribute: "DARK",
				Type: "Spellcaster",
			},
			{
				Name: "Red-Eyes Black Dragon",
				Description: "A ferocious dragon with a deadly attack.",
				Attack: 2400,
				Defense: 2000,
				Level: 7,
				Attribute: "DARK",
				Type: "Dragon",
			},
		}
		c.JSON(http.StatusOK, cards)
	})

	// WARNING - This should be removed after first run, it's only to get all cards from the API and save them to the DB
	r.GET("/cards/getcards", func(c *gin.Context) {
		allCards := conn.GetCards()
		conn.SaveCards(allCards)
		c.JSON(http.StatusOK, "{\"message\": \"Cards saved to DB\"}")
	})

	r.GET("/cards/getcardfromdb", func(c *gin.Context) {
		card := conn.FetchCard("Blue-Eyes White Dragon")
		c.JSON(http.StatusOK, card)
	})
	
	r.Run(":8080")
}

// Simple function to fetch all tables from the DB
// NOTE: this is just a test function, you should remove it once you understand how it works
func FetchTables() []string {
	query, err := conn.DB.Query("SHOW TABLES")
	if err != nil {
		panic("SQL ERROR")
	}

	tables := []string{}

	println("Tables:")
	for query.Next() {
		var tableName string
		err = query.Scan(&tableName)
		if err != nil {
			panic("SQL ERROR")
		}
		tables = append(tables, tableName)
	}

	return tables
}

// Path: main.go