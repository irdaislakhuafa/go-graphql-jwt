package directives

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/irdaislakhuafa/go-graphql-jwt/middlewares"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func Auth(ctx context.Context, object interface{}, next graphql.Resolver) (interface{}, error) {
	tokenClaims := middlewares.CtxValue(ctx)
	if tokenClaims == nil {
		return nil, &gqlerror.Error{
			Message: "Access denied",
		}
	}

	return next(ctx)
}
