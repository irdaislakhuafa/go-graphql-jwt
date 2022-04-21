package middlewares

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/irdaislakhuafa/go-graphql-jwt/service"
)

type authString string

func AuthMiddleware(next http.Handler) http.Handler {
	log.Println("Entering middleware")
	handlerFunc := func(writer http.ResponseWriter, request *http.Request) {
		log.Println("Get Authorization field from header")
		auth := request.Header.Get("Authorization")
		if strings.Trim(auth, " ") == "" {
			log.Println("Authorization field is empty!")
			next.ServeHTTP(writer, request)
			return
		}

		log.Println("Get JWT token from Bearer")
		bearer := "Bearer "
		token := auth[len(bearer):]

		log.Println("Validating token")
		jwtToken, err := service.JwtValidate(context.Background(), token)
		if err != nil || !jwtToken.Valid {
			log.Println("Invalid token!")
			http.Error(writer, "Invalid token", http.StatusForbidden)
			return
		}
		log.Println("Token is valid")

		log.Println("Calims payload from token")
		claims, _ := jwtToken.Claims.(*service.JwtCustomClaims)

		log.Println("Create context with key and value claims")
		ctx := context.WithValue(request.Context(), authString("auth"), claims)

		log.Println("Overwrite request with new context")
		request = request.WithContext(ctx)

		log.Println("Continue to next service")
		next.ServeHTTP(writer, request)
	}

	return http.HandlerFunc(handlerFunc)
}

func CtxValue(ctx context.Context) *service.JwtCustomClaims {
	log.Println("Get claims from context")
	value, _ := ctx.Value(authString("auth")).(*service.JwtCustomClaims)
	return value
}
