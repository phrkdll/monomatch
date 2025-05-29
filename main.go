package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/phrkdll/monomatch/api"
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

	api.RegisterRoutes(r)

	return r
}
