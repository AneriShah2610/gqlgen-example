package main

import (
	log "log"
	http "net/http"
	os "os"

	handler "github.com/99designs/gqlgen/handler"
	gqlgen_example "github.com/aneri/gqlgen-example"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", gqlgen_example.MiddleWareHandler(handler.GraphQL(gqlgen_example.NewExecutableSchema(gqlgen_example.Config{Resolvers: &gqlgen_example.Resolver{}}))))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
