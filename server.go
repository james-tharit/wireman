package main

import (
	"log"
	"net/http"
	"os"
	"wireman/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

const defaultPort = "8080"

func main() {

	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}

	router := chi.NewRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	graphqlServer := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	router.Handle("/", playground.Handler("Wireman playground", "/query"))
	router.Handle("/query", graphqlServer)

	log.Printf("connect to http://localhost:%s/ for Wireman playground", port)
	serveErr := http.ListenAndServe(":"+port, router)
	if serveErr != nil {
		log.Fatal(http.ListenAndServe(":"+port, nil))
	}
}
