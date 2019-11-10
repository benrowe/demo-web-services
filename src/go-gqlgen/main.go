package main

import (
    "context"
    "fmt"
    "github.com/99designs/gqlgen/graphql"
    "github.com/99designs/gqlgen/handler"
    "github.com/benrowe/demo-web-services/src/app"
    "github.com/benrowe/demo-web-services/src/app/models"
    "github.com/benrowe/demo-web-services/src/go-gqlgen/gen"
    "github.com/benrowe/demo-web-services/src/go-gqlgen/resolvers"
    _ "github.com/go-sql-driver/mysql"
    "github.com/gorilla/mux"
    "github.com/thedevsaddam/govalidator"

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
    c := gen.Config{Resolvers: &resolvers.Resolver{
        App: app,
    }}
    loadDirectives(&c.Directives)

    r.Handle("/query", handler.GraphQL(gen.NewExecutableSchema(c)))

    log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
    if err := http.ListenAndServe(":"+port, r); err != nil {
        log.Fatalf("unable to serve request: %v", err)
    }
}

func loadDirectives(d *gen.DirectiveRoot) {
    d.Constraint = func(ctx context.Context, obj interface{}, next graphql.Resolver, rules []string) (res interface{}, err error) {

        log.Printf("validating: %v\n", rules)
        type value struct{ v interface{} }
        r := govalidator.MapData{
            "v": rules,
        }
        opt := govalidator.Options{
            Rules: r,
            Data:  &value{v: obj},
        }
        v := govalidator.New(opt)
        e := v.ValidateStruct()
        log.Print(e)
        if len(e) > 0 {
            log.Printf("%+v", e)
            return nil, fmt.Errorf("%s", e)
        }

        return next(ctx)
    }
}

func printRequest(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("receiving request: %v", r.URL.Path)
        next.ServeHTTP(w, r)
    })
}
