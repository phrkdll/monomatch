package session

import (
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(mux *chi.Mux) {
	mux.Route("/sessions", func(r chi.Router) {
		r.Post("/", createSession)
		r.Get("/{sessionId}", getSessionInfo)
		r.Post("/{sessionId}/join", joinSession)
	})
}
