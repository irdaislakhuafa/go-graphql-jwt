package service

import (
	"context"
	"log"
	"time"

	"gopkg.in/dgrijalva/jwt-go.v3"
)

type JwtCustomClaims struct {
	ID string `json:"id"`
	jwt.StandardClaims
}

var keySecret = []byte("MySecret")

var (
	expiredAt int64 = time.Now().Add(time.Hour * 72).Unix()
	issueAt   int64 = time.Now().Unix() // createdAt
)

func JwtGenerate(ctx context.Context, userID string) (string, error) {
	log.Println("Initilaized claims for JWT")
	jwtClaims := &JwtCustomClaims{
		ID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredAt,
			IssuedAt:  issueAt,
		},
	}

	log.Println("Generate new JWT with algorithn HS256 and custom claims")
	jwt := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)

	log.Println("Signed key secret string to token")
	token, err := jwt.SignedString(keySecret)
	if err != nil {
		log.Println("Error while generate token:", err.Error())
		return "", err
	}

	return token, nil
}
