package main

import (
    "github.com/99designs/gqlgen/handler"
    "github.com/gorilla/mux"
    "github.com/benrowe/demo-web-services/src/go-gqlgen/gen"
    "github.com/benrowe/demo-web-services/src/go-gqlgen/resolvers"
    "log"
    "net/http"
    "os"
)

const defaultPort = "8888"

func main() {

    port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

    //http.Handle("/query", handler.GraphQL(NewExecutableSchema(Config{})))
	//http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	//http.Handle("/query", handler.GraphQL(NewExecutableSchema(Config{Resolvers: &Resolver{}})))

    r := mux.NewRouter()

    r.Use(printRequest)
    r.Handle("/", handler.Playground("test", "/query"))
    r.Handle("/query", handler.GraphQL(gen.NewExecutableSchema(gen.Config{Resolvers: &resolvers.Resolver{}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
	    log.Fatalf("unable to serve request: %v", err)
    }
}

func printRequest(next http.Handler) http.Handler {
    return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
        log.Printf("receiving request: %v", r.URL.Path)
        next.ServeHTTP(w, r)
    })
}

