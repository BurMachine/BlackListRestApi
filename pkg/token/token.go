package token

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

func GenerateToken(name string) (string, error) {
	claims := jwt.MapClaims{}
	claims["sub"] = uuid.New()
	claims["name"] = name

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := "mysecretkey"

	return token.SignedString([]byte(secret))
}
