package user

import (
	// conn "github.com/Aesir-Development/yugioh-backend/internal/db"
)

// User struct
type User struct {
	ID int64
	PUUID string
	Username string
	Password string
}

func ParseUsers(body []byte) []User {
	// TODO - Parse the users from the body
	return []User{}
}