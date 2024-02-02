package login

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var secretKey = []byte("secret")

// CreateToken creates a new token
func CreateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims {
			"username": username,
			"exp": time.Now().Add(time.Hour * 24).Unix(),
		})
	
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", fmt.Errorf("error creating token")
	}
	
	return tokenString, nil
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return fmt.Errorf("error parsing token")
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}