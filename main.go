package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	conn "github.com/Aesir-Development/yugioh-backend/internal/db" // Importing the DB connection package
	"github.com/Aesir-Development/yugioh-backend/pkg/card" // Importing the card package
)


func main() {
	r := gin.Default()
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

	// TODO - Make a better test route for the card package. This is just a placeholder
	r.GET("/test", func(c *gin.Context) {
		card.TestCardFetch("Blue-Eyes White Dragon")
		tables := FetchTables()
		c.JSON(http.StatusOK, tables)
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