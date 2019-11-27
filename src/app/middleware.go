package app

import (
	"log"
	"net/http"
)

func PrintRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("receiving request: %v", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
