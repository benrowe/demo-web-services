package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("content-type", "application/json")
		_, err := w.Write([]byte(`["OK"]`))
		if err != nil {
			log.Printf("unable to write bytes to response: %v", err)
		}
	}).Methods("GET")

	r.Use(printRequest)

	log.Printf("serving on :80")
	if err := http.ListenAndServe(":80", r); err != nil {
		log.Fatalf("unable to serve request: %v", err)
	}
}

func printRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("receiving request: %v", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
