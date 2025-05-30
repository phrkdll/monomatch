package session

import (
	"encoding/json"
	"net/http"

	"github.com/phrkdll/monomatch/pkg/session"
	"github.com/phrkdll/monomatch/pkg/session/store"
)

type JoinRequest struct {
	Id         session.SessionId `json:"id"`
	PlayerName string            `json:"playerName"`
}

func joinSession(w http.ResponseWriter, r *http.Request) {
	var request JoinRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))

		return
	}

	session, err := store.Instance().Get(request.Id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))

		return
	}

	p, err := session.AddPlayer(request.PlayerName)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))

		return
	}

	json, err := json.Marshal(&p)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(json)
}
