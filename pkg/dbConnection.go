package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Global DB context

var DB *sql.DB

func init() {
	println("DB Init")

	var err error
	DB, err = sql.Open("mysql", "DB HERE") // TODO add DB connection string to an env.
	if err != nil {
		panic("FUUUCK")
	}

	err = DB.Ping()
	if err != nil {
		panic(err)
	}
}