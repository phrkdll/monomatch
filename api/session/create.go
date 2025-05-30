package session

import (
	"encoding/json"
	"net/http"

	"github.com/phrkdll/monomatch/pkg/session"
	"github.com/phrkdll/monomatch/pkg/session/store"
)

type CreateSessionRequest struct {
	SessionName string   `json:"sessionName"`
	Symbols     []string `json:"symbols"`
}

func createSession(w http.ResponseWriter, r *http.Request) {
	var request CreateSessionRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))

		return
	}

	session, err := session.New(request.SessionName, request.Symbols)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))

		return
	}

	store.Instance().Add(session)

	json, err := json.Marshal(&session)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(json)
}
