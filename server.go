package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gographql-project/graph"
	"github.com/gographql-project/graph/generated"
)
// steps to build apps using graphql
// step-1 setup graphql project
// step-2 write schema on schema.graphqls
// step-3 write query (implement a Query that we defined in schema.graphqls.) go to schema.resolver.go and define
// functions which you have declare in Mutations
// step -4 define your functions  in schema.resolvers.go and delete generate.go file
// step-5 go run github.com/99designs/gqlgen generate (to get generate.go file)
// if not work then try -> go generate ./...

// go run server.go and check
// step-6 for query -->
// query {
//	links{
//    title
//    address,
//    user{
//      name
//    }
//  }
//}
// for mutations
//mutation {
//createLink(input: {title: "new link", address:"http://address.org"}){
//title,
//user{
//name
//}
//address
//}
//}
const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
