package session

import (
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(mux *chi.Mux) {
	mux.Route("/api/sessions", func(r chi.Router) {
		r.Post("/", createSession)
	})

	mux.Route("/ws/sessions", func(r chi.Router) {
		r.HandleFunc("/", sessionSocket)
	})
}
