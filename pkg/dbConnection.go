package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func db(name string) {
	// db connection
	db, _ := sql.Open("mysql", "URL HERE")
	defer db.Close()

	version := ""

	db.QueryRow("SELECT VERSION()").Scan(&version)
	fmt.Println("Connected to:", version)
}
