package model

import "github.com/golang-jwt/jwt/v4"

type TokenClaim struct {
	Email   string `json:"email"`
	IsAdmin bool   `json:"is_admin"`
	jwt.RegisteredClaims
}
