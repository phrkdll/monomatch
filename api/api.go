package api

import "github.com/go-chi/chi/v5"

func RegisterRoutes(mux *chi.Mux) {
	registerSessionRoutes(mux)
}
