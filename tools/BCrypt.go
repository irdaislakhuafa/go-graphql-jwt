package tools

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func EncodePassword(password string) (string, error) {
	log.Println("Ecrypting password")
	encode, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Failed when encrypt password:", err.Error())
		return "", err
	}
	log.Println("Success encrypt password")
	return string(encode), nil
}

func VerifyPassword(encodedPassword string, realPassword string) error {
	errorStatus := bcrypt.CompareHashAndPassword([]byte(encodedPassword), []byte(realPassword))
	return errorStatus
}
