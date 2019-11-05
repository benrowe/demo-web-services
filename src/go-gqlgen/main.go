package main

import (
    "github.com/99designs/gqlgen/handler"
    "github.com/benrowe/demo-web-services/src/app"
    "github.com/gorilla/mux"
    "github.com/benrowe/demo-web-services/src/go-gqlgen/gen"
    "github.com/benrowe/demo-web-services/src/go-gqlgen/resolvers"
    "github.com/jinzhu/gorm"
    _ "github.com/go-sql-driver/mysql"

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
    log.Printf("%+v", app)

    port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

    db, err := gorm.Open("mysql", "root:password@(database)/employees?charset=utf8&parseTime=True&loc=Local")
    if err != nil {
        log.Fatalf("Unable to connect to database: %v", err)
    }

    defer db.Close()


    r := mux.NewRouter()

    // middleware
    r.Use(printRequest)

    // routes
    r.Handle("/", handler.Playground("test", "/query"))
    r.Handle("/query", handler.GraphQL(gen.NewExecutableSchema(gen.Config{Resolvers: &resolvers.Resolver{
        DB: db,
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

