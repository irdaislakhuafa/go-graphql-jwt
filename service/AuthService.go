package service

import (
	"context"
	"log"

	"github.com/irdaislakhuafa/go-graphql-jwt/graph/model"
	"github.com/irdaislakhuafa/go-graphql-jwt/tools"
	"gorm.io/gorm"
)

func UserRegister(ctx context.Context, newUser model.NewUser) (interface{}, error) {
	_, err := UserGetByEmail(ctx, newUser.Email)
	if err == gorm.ErrRecordNotFound {
		return nil, err
	} else if err != nil {
		return nil, err
	}

	createdUser, err := UserCreate(ctx, newUser)
	if err != nil {
		return nil, err
	}

	token, err := JwtGenerate(ctx, createdUser.ID)
	if err != nil {
		return nil, err
	}

	response := map[string]interface{}{
		"token": token,
	}

	return response, nil
}

func UserLogin(ctx context.Context, email string, password string) (interface{}, error) {
	user, err := UserGetByEmail(ctx, email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Println("User with email:", email, "not found")
			return nil, err
		}
		return nil, err
	}

	log.Println("Verify password")
	if err := tools.VerifyPassword(user.Password, password); err != nil {
		log.Println("Password not valid")
		return nil, err
	}
	generatedToken, err := JwtGenerate(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	response := map[string]interface{}{
		"token": generatedToken,
	}

	return response, nil
}
