package session

import (
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(mux *chi.Mux) {
	mux.Route("/session", func(r chi.Router) {
		r.Post("/create", createSession)
		r.Post("/join", joinSession)
	})
}
