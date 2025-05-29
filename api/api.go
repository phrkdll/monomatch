package api

import (
	"github.com/go-chi/chi/v5"

	"github.com/phrkdll/monomatch/api/session"
)

func RegisterRoutes(mux *chi.Mux) {
	session.RegisterRoutes(mux)
}
