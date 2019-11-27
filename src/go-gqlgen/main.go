package main

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
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

	a := app.NewApp(cfg)

	models.Migrate(a.DB)

	defer a.DB.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	r := mux.NewRouter()

	// middleware
	r.Use(app.PrintRequest)

	// routes
	r.Handle("/", handler.Playground("test", "/query"))
	c := gen.Config{Resolvers: &resolvers.Resolver{
		App: a,
	}}

	loadDirectives(&c.Directives)

	r.Handle("/query", handler.GraphQL(gen.NewExecutableSchema(c)))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("unable to serve request: %v", err)
	}
}

func loadDirectives(d *gen.DirectiveRoot) {
	d.Constraint = func(ctx context.Context, obj interface{}, next graphql.Resolver, name string, rules []string) (res interface{}, err error) {
		// convert rules
		if msg, ok := app.ValidateVal(obj, name, convertRules(rules)); !ok {
			return nil, fmt.Errorf("invalid input: %s", msg)
		}
		return next(ctx)
	}
}

func convertRules(rules []string) *[]app.ValidationRule {
	newRules := []app.ValidationRule{}

	for i, rule := range rules {
		newRules[i] = app.ValidationRule(rule)
	}

	return &newRules
}
