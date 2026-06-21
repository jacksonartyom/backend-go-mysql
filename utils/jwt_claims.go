package utils

import "github.com/golang-jwt/jwt/v5"

type JwtClaims struct {
	UserId string `json:"userId"`
	jwt.RegisteredClaims
}
