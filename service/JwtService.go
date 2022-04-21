package service

import "gopkg.in/dgrijalva/jwt-go.v3"

type JwtCustomClaims struct {
	ID string `json:"id"`
	jwt.StandardClaims
}
