package login

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/Aesir-Development/yugioh-backend/internal/user"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var u user.User

	json.NewDecoder(r.Body).Decode(&u)
	fmt.Printf("User: %v\n", u)

	if u.Username == "test" && u.Password == "test" {
		token, err := CreateToken(u.Username)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error": "error creating token"}`))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"token": "` + token + `"}`))

		fmt.Printf("Token: %v\n", token)

		return
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"error": "invalid credentials"}`))
		return
	}
}