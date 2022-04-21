package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/irdaislakhuafa/go-graphql-jwt/config"
	"github.com/irdaislakhuafa/go-graphql-jwt/entity"
	"github.com/irdaislakhuafa/go-graphql-jwt/graph/model"
	"github.com/irdaislakhuafa/go-graphql-jwt/tools"
)

func UserCreate(ctx context.Context, newUser model.NewUser) (*entity.User, error) {
	var err error
	db := config.GetDB()

	newUser.Password, err = tools.EncodePassword(newUser.Password)
	if err != nil {
		return nil, err
	}

	user := entity.User{
		ID:       uuid.New().String(),
		Name:     newUser.Name,
		Email:    newUser.Email,
		Password: newUser.Password,
	}

	if err := db.Model(user).Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// func UserGetById(ctx context.Context, id string) ()
