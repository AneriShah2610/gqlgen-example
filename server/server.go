package main

import (
	log "log"
	http "net/http"
	os "os"

	handler "github.com/99designs/gqlgen/handler"
	gqlgen_example "github.com/aneri/gqlgen-example"
	"github.com/gorilla/mux"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	router := mux.NewRouter()
	router.Handle("/", handler.Playground("GraphQL playground", "/query"))
	router.Handle("/query", gqlgen_example.MiddleWareHandler(handler.GraphQL(gqlgen_example.NewExecutableSchema(gqlgen_example.Config{Resolvers: &gqlgen_example.Resolver{}}))))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
