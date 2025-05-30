package main

import (
	"github.com/go-chi/chi/v5"

	"github.com/phrkdll/monomatch/api/session"
)

func registerRoutes(mux *chi.Mux) {
	session.RegisterRoutes(mux)
}
