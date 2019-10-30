package main

import (
    "fmt"
    "net/http"
)

type service struct {
    container string
    port string
}

func (s service) url() string {
    return fmt.Sprintf("http://%v:%v", "localhost", s.port)
}

func main() {

    services := []*service{
        &service{container:"go-grpc", port:"802"},
        &service{container:"go-gqlgen", port: "801"},
        &service{container:"go-swagger", port:"803"},
    }

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("content-type", "text/html")
        // generate a list of implementations
        w.Write([]byte("<ul>"))
        for _, s := range services {
            w.Write([]byte("<li><a href=\""+s.url()+"\">"+s.container+"</a></li>"))
        }
        w.Write([]byte("</ul>"))

    })
    http.ListenAndServe(":80", nil)
}
