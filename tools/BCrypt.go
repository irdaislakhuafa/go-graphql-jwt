package tools

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) (string, error) {
	log.Println("Ecrypting password")
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Failed when encrypt password:", err.Error())
		return "", err
	}
	log.Println("Success encrypt password")
	return string(encryptedPassword), nil
}
