package util

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWTClaim struct {
	Id    string
	Email string
	jwt.RegisteredClaims
}

func NewJwtClaim(id string, email string, issuer string, expTime time.Time) *JWTClaim {
	return &JWTClaim{
		Id:    id,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    issuer,
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}
}

func (j *JWTClaim) SignToken(jwtKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, j)
	signedToken, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return signedToken, nil
}
