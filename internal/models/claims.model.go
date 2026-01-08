package models

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/jerson2000/jquest/internal/enums"
)

type JWTClaims struct {
	Id   int        `json:"id"`
	Name string     `json:"name"`
	Role enums.Role `json:"role"`
	jwt.RegisteredClaims
}
