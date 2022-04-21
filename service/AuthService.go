package service

import (
	"context"

	"github.com/irdaislakhuafa/go-graphql-jwt/graph/model"
	"gorm.io/gorm"
)

func UserRegister(ctx context.Context, newUser model.NewUser) (interface{}, error) {
	_, err := UserGetByEmail(ctx, newUser.Email)
	if err == gorm.ErrRecordNotFound {
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
