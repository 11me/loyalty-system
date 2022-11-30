package model

import "github.com/golang-jwt/jwt/v4"

type TokenClaim struct {
	Sub string `json:"sub"`
	jwt.RegisteredClaims
}
