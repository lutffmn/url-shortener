package main

import (
	"log"
	"net/http"

	h "github.com/lutffmn/url-shortener/internal/handlers"
)

func main() {
	r := http.NewServeMux()

	r.HandleFunc("POST /api/v1/shorten", h.ShortenHandler)
	r.HandleFunc("GET /api/v1/{alias}", func(w http.ResponseWriter, r *http.Request) {})
	r.HandleFunc("GET /api/v1/urls/{alias}", func(w http.ResponseWriter, r *http.Request) {})
	r.HandleFunc("GET /api/v1/urls", func(w http.ResponseWriter, r *http.Request) {})
	r.HandleFunc("GET /api/v1/urls/{alias}/stats", func(w http.ResponseWriter, r *http.Request) {})
	r.HandleFunc("PATCH /api/v1/urls/{alias}", func(w http.ResponseWriter, r *http.Request) {})
	r.HandleFunc("DELETE /api/v1/urls/{alias}", func(w http.ResponseWriter, r *http.Request) {})
	r.HandleFunc("GET /api/v1/health", h.HealthHandler)

	log.Fatal(http.ListenAndServe(":8080", r))
}
