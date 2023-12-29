package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/nqnlong1506/cqrs-event-sourcing-go/pkg/graph"
	"github.com/nqnlong1506/cqrs-event-sourcing-go/pkg/mongo"
	repo "github.com/nqnlong1506/cqrs-event-sourcing-go/pkg/repository"
)

const defaultPort = "8080"

func main() {
	// envFile, _ := godotenv.Read(".env")
	// fmt.Println(envFile["http_ca_cert"])

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func init() {
	mongo.ConnectDB()
	repo.InitBookCollection()
	repo.InitTransactionCollection()
	repo.InitUserCollection()
}
