package main

import (
    "github.com/99designs/gqlgen/handler"
    "github.com/benrowe/demo-web-services/src/app"
    "github.com/benrowe/demo-web-services/src/app/models"
    "github.com/benrowe/demo-web-services/src/go-gqlgen/gen"
    "github.com/benrowe/demo-web-services/src/go-gqlgen/resolvers"
    _ "github.com/go-sql-driver/mysql"
    "github.com/gorilla/mux"

    "log"
    "net/http"
    "os"
)

const defaultPort = "8888"

type env struct{}

func (e env) LookupEnv(key string) (string, bool) {
    return os.LookupEnv(key)
}

func main() {

    cfg := app.LoadConfigFromEnv(&env{})

    app := app.NewApp(cfg)

    models.Migrate(app.DB)

    defer app.DB.Close()

    log.Printf("%+v", app)

    port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

    r := mux.NewRouter()

    // middleware
    r.Use(printRequest)

    // routes
    r.Handle("/", handler.Playground("test", "/query"))
    r.Handle("/query", handler.GraphQL(gen.NewExecutableSchema(gen.Config{Resolvers: &resolvers.Resolver{
        App: app,
    }})))

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

