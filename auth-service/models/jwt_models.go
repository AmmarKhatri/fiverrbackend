package models

import "github.com/golang-jwt/jwt/v4"

type Claims struct {
	ID int64 `json:"id"`
	jwt.RegisteredClaims
}

type ExternalClaims struct {
	ID int64 `json:"id"`
	jwt.RegisteredClaims
}
