package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/phrkdll/monomatch/pkg/models/session"
)

func registerSessionRoutes(mux *chi.Mux) {
	mux.Route("/sessions", func(r chi.Router) {
		r.Post("/new", createNewSession)
	})
}

func createNewSession(w http.ResponseWriter, r *http.Request) {
	var input []string

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))

		return
	}

	session, err := session.New(input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))

		return
	}

	json, err := json.Marshal(&session)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(json)
}
