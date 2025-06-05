package session

import (
	"encoding/json"
	"net/http"

	"github.com/phrkdll/monomatch/pkg/session"
	"github.com/phrkdll/monomatch/pkg/session/store"
	"github.com/phrkdll/must/pkg/must"
)

type CreateSessionRequest struct {
	SessionName string   `json:"sessionName"`
	Symbols     []string `json:"symbols"`
}

func createSession(w http.ResponseWriter, r *http.Request) {
	defer must.Recover()

	var request CreateSessionRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	must.SucceedOr(err).Respond(w, http.StatusBadRequest)
	must.Succeed(err)

	session, err := session.New(request.SessionName, request.Symbols)
	must.SucceedOr(err).Respond(w, http.StatusBadRequest)
	must.Succeed(err)

	store.Instance().Add(session)

	json, err := json.Marshal(&session)
	must.SucceedOr(err).Respond(w, http.StatusBadRequest)
	must.Succeed(err)

	w.WriteHeader(http.StatusOK)
	w.Write(json)
}
