package common

import (
	"crypto/sha1"
	"encoding/hex"

	"github.com/dgrijalva/jwt-go"
)

type MyCustomClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func JWT(priKey []byte, username string, tokenExpire int64) string {
	claims := MyCustomClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(tokenExpire)).Unix(),
			Issuer:    "vpn",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, _ := token.SignedString(priKey)
	return tokenString
}

func ParseJWT(priKey []byte, token string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return priKey, nil
	})
}
