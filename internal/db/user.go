package dbConnection

import (
	"net/http"
	"io"
	"github.com/Aesir-Development/yugioh-backend/internal/user"
	"encoding/json"
)


func GetUsers() []user.User {
	resp, err := http.Get("https://db.ygoprodeck.com/api/v7/cardinfo.php")
	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic("Error reading response body")
	}

	users := user.ParseUsers(body)

	return users
}

func SaveUsers(users []user.User) {
	for _, user := range users {
		_, err := DB.Exec("INSERT INTO users (p_uuid, username, password) VALUES (?, ?, ?)",
		user.PUUID, user.Username, user.Password)
	
		if err != nil {
			panic(err)
		}
	}
}

func UserToJSON(user user.User) string {
	userJSON, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	return string(userJSON)
}

func UsersToJSON(users []user.User) string {
	usersJSON, err := json.Marshal(users)
	if err != nil {
		panic(err)
	}

	return string(usersJSON)
}

func JSONToUser(userJSON string) user.User {
	var user user.User

	err := json.Unmarshal([]byte(userJSON), &user)
	if err != nil {
		panic(err)
	}

	return user
}
