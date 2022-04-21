package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/irdaislakhuafa/go-graphql-jwt/graph/generated"
	"github.com/irdaislakhuafa/go-graphql-jwt/graph/model"
	"github.com/irdaislakhuafa/go-graphql-jwt/service"
)

func (r *authResolver) Login(ctx context.Context, obj *model.Auth, email string, password string) (interface{}, error) {
	return service.UserLogin(ctx, email, password)
}

func (r *authResolver) Register(ctx context.Context, obj *model.Auth, newUser model.NewUser) (interface{}, error) {
	return service.UserRegister(ctx, newUser)
}

func (r *mutationResolver) Auth(ctx context.Context) (*model.Auth, error) {
	return &model.Auth{}, nil
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	user, err := service.UserGetById(ctx, id)
	modelUser := model.User{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}

	return &modelUser, err

}

// Auth returns generated.AuthResolver implementation.
func (r *Resolver) Auth() generated.AuthResolver { return &authResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type authResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
