package session

import (
	"encoding/json"
	"net/http"

	"github.com/phrkdll/monomatch/pkg/session"
	"github.com/phrkdll/monomatch/pkg/session/store"
	"github.com/phrkdll/must/pkg/must"
)

type JoinRequest struct {
	Id         session.SessionId `json:"id"`
	PlayerName string            `json:"playerName"`
}

func joinSession(w http.ResponseWriter, r *http.Request) {
	defer must.Recover()

	var request JoinRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	must.SucceedOr(err).Respond(w, http.StatusBadRequest)
	must.Succeed(err)

	session, err := store.Instance().Get(request.Id)
	must.SucceedOr(err).Respond(w, http.StatusBadRequest)
	must.Succeed(err)

	p, err := session.AddPlayer(request.PlayerName)
	must.SucceedOr(err).Respond(w, http.StatusBadRequest)
	must.Succeed(err)

	json, err := json.Marshal(&p)
	must.SucceedOr(err).Respond(w, http.StatusBadRequest)
	must.Succeed(err)

	w.WriteHeader(http.StatusOK)
	w.Write(json)
}
