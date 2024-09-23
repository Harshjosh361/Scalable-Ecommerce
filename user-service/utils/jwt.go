package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	UserID string `json:"_id"`
	jwt.StandardClaims
}

var jwtSecret = []byte("your_secret_key")

func GenerateJWT(userID string) (string, error) {
	expirationTime := time.Now().Add(time.Hour * 24)

	// creating claims
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ValidateJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}

	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid or expired token")
	}

	return claims, nil
}
