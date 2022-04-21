package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/irdaislakhuafa/go-graphql-jwt/service"
)

type authString string

func AuthMiddleware(next http.Handler) http.Handler {
	handlerFunc := func(writer http.ResponseWriter, request *http.Request) {
		auth := request.Header.Get("Authorization")
		if strings.Trim(auth, " ") == "" {
			next.ServeHTTP(writer, request)
			return
		}

		bearer := "Bearer "
		token := auth[len(bearer):]

		validate, err := service.JwtValidate(context.Background(), token)
		if err != nil || !validate.Valid {
			http.Error(writer, "Invalid token", http.StatusForbidden)
			return
		}

		claims, _ := validate.Claims.(*service.JwtCustomClaims)
		ctx := context.WithValue(request.Context(), authString("auth"), claims)
		request = request.WithContext(ctx)
		next.ServeHTTP(writer, request)
	}

	return http.HandlerFunc(handlerFunc)
}

func CtxValue(ctx context.Context) *service.JwtCustomClaims {
	value, _ := ctx.Value(authString("auth")).(*service.JwtCustomClaims)
	return value
}
