package utils

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
	Id        uint   `json:"id"`
	Username  string `json:"username"`
	Authority int    `json:"authority"`
	jwt.StandardClaims
}

func GenerateUserToken(id uint, username string, authority int) (string, error) {
	nomTime := time.Now()
	expireTime := nomTime.Add(24 * time.Hour)
	claims := Claims{
		Id:        id,
		Username:  username,
		Authority: authority,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "test-todo-list",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
		return claims, nil
	}

	return nil, err
}
