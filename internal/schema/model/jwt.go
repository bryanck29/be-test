package model

import (
	"github.com/dgrijalva/jwt-go"
)

// JwtClaims represent JWT claims
type JwtClaims struct {
	User User `json:"user"`
	jwt.StandardClaims
}
