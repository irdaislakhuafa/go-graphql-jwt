package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
	"github.com/irdaislakhuafa/go-graphql-jwt/config"
	"github.com/irdaislakhuafa/go-graphql-jwt/directives"
	"github.com/irdaislakhuafa/go-graphql-jwt/graph"
	"github.com/irdaislakhuafa/go-graphql-jwt/graph/generated"
	"github.com/irdaislakhuafa/go-graphql-jwt/middlewares"
	"github.com/irdaislakhuafa/go-graphql-jwt/migration"
)

func main() {
	options := config.Config{}
	options.EnableFlags(true)

	database := config.Database{}
	database.Init()

	migration.EnableMigration(true)

	router := mux.NewRouter()
	router.Use(middlewares.AuthMiddleware)

	gqlConfig := generated.Config{Resolvers: &graph.Resolver{}}
	gqlConfig.Directives.Auth = directives.Auth

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(gqlConfig))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", options.ServerPort)
	log.Fatal(http.ListenAndServe(":"+options.ServerPort, router))
}
