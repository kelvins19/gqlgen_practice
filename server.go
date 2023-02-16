package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	generated "github.com/kelvins19/BCX_BE/graph"
	graph "github.com/kelvins19/BCX_BE/graph"
	"github.com/spf13/viper"
)

const defaultPort = "8080"

func main() {
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	// Set a default value for the port number
	viper.SetDefault("PORT", "8080")
	port := viper.GetString("PORT")
	fmt.Printf("Starting server on port %s\n", port)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Database := graph.Connect()
	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: Database}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
