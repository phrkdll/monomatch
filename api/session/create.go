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

	must.Succeed(json.NewDecoder(r.Body).Decode(&request)).
		ElseRespond(w, http.StatusBadRequest).
		ElsePanic()

	session := must.Return(session.New(request.SessionName, request.Symbols)).
		ElseRespond(w, http.StatusBadRequest).
		ElsePanic()

	store.Instance().Add(session)

	json := must.Return(json.Marshal(&session)).
		ElseRespond(w, http.StatusBadRequest).
		ElsePanic()

	w.WriteHeader(http.StatusOK)
	w.Write(json)
}
