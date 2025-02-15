package auth

import (
	"github.com/golang-jwt/jwt/v4"
)

var JWT_SECRET = "secret"

func JWT(text string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"text": text,
	})

	return token.SignedString([]byte(JWT_SECRET))
}
