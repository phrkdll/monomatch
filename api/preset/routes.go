package preset

import (
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(mux *chi.Mux) {
	mux.Route("/presets", func(r chi.Router) {
		r.Get("/", listPresets)
	})
}
