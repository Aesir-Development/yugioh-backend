package dbConnection

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"os"
)

// Global DB context

var DB *sql.DB

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	println("DB Init")

	// Loading the DB connection string from the .env file
	conString := os.Getenv("DB_CONNECTION_STRING")

	DB, err = sql.Open("mysql", conString)
	if err != nil {
		panic(err)
	}

	err = DB.Ping()
	if err != nil {
		panic(err)
	}
}