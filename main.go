package main

import (
	"fmt"
	hydra "github.com/ory/hydra/sdk/go/hydra/client"
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type service struct {
	Container string
	port      string
}

func (s service) url() string {
	return fmt.Sprintf("http://%v:%v", "localhost", s.port)
}

func loadServices() []*service {
	return []*service{
		&service{Container: "go-grpc", port: "802"},
		&service{Container: "go-gqlgen", port: "801"},
		&service{Container: "go-swagger", port: "803"},
	}
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "text/html")
		t := template.New("services.html.tmpl")
		t, err := t.ParseFiles("./web/template/services.html.tmpl")

		if err != nil {
			log.Printf("unable to parse template: %v", err)
		}

		err = t.Execute(w, loadServices())
		if err != nil {
			log.Println(err)
		}
	})

	http.HandleFunc("/auth/logout", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("not implemented: logout"))
	})

	http.HandleFunc("/auth/login", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("not implemented: login"))

	})

	http.HandleFunc("/auth/consent", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("not implemented: consent"))
	})

	http.HandleFunc("/.well-known/oauth-authorization-server", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
	})

	http.ListenAndServe(":80", nil)
}

func getHyrda() *hydra.OryHydra {
	u, _ := url.Parse("http://127.0.0.1:4444")
	return hydra.NewHTTPClientWithConfig(nil, &hydra.TransportConfig{
		Schemes: []string{
			u.Scheme,
		},
		Host:     u.Host,
		BasePath: u.Path,
	})
}
