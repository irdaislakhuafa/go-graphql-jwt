package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/irdaislakhuafa/go-graphql-jwt/config"
	"github.com/irdaislakhuafa/go-graphql-jwt/graph"
	"github.com/irdaislakhuafa/go-graphql-jwt/graph/generated"
)

func main() {
	options := config.Config{}
	options.EnableFlags(true)

	database := config.Database{}
	database.Init()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", options.ServerPort)
	log.Fatal(http.ListenAndServe(":"+options.ServerPort, nil))
}
