package dto

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTClaims struct {
	UserId int    `json:"user_id"`
	Email  string `json:"email"`
	RoleId int    `json:"role_id"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(claims JWTClaims) (string, error) {
	expirationTime := time.Now().Add(time.Hour * 24 * 30)
	claims.ExpiresAt = jwt.NewNumericDate(expirationTime)
	claims.Issuer = os.Getenv("APP_NAME")
	claims.IssuedAt = jwt.NewNumericDate(time.Now())
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("API_SECRET")))
}
