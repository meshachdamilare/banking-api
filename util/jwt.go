package util

import (
	"errors"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/meshachdamilare/banking-api/config"
	"time"
)

type JWTClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

var (
	issuer    = config.GetEnv("ISSUER")
	secretKey = []byte(config.GetEnv("SECRET_KEY"))
)

func GenerateToken(email string) (string, error) {
	claims := JWTClaims{
		email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
			Issuer:    issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func ParseToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
