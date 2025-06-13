package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/phrkdll/monomatch/api/preset"
	"github.com/phrkdll/monomatch/api/session"
)

func main() {
	host := "localhost:1982"

	fmt.Printf("[MMA] Starting\n")

	// initDatabase()
	r := initChi()

	fmt.Printf("[MMA] Listening on %s\n", host)

	panic(http.ListenAndServe(host, r))
}

func initChi() *chi.Mux {
	fmt.Printf("[MMA] Setting up routes\n")

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	registerRoutes(r)

	return r
}

func registerRoutes(mux *chi.Mux) {
	session.RegisterRoutes(mux)
	preset.RegisterRoutes(mux)
}
