package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	conn "github.com/Aesir-Development/yugioh-backend/internal/db" // Importing the DB connection package
)

// Card is a simple example struct
// TODO - modify this to match your needs
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

	// TODO - remove this test route
	r.GET("/test", func(c *gin.Context) {
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