package session

import (
	"encoding/json"
	"net/http"

	"github.com/phrkdll/monomatch/pkg/preset"
	"github.com/phrkdll/monomatch/pkg/session"
	"github.com/phrkdll/monomatch/pkg/session/store"
	"github.com/phrkdll/must/pkg/must"
)

type CreateSessionRequest struct {
	SessionName string           `json:"sessionName"`
	Preset      *preset.PresetId `json:"preset"`
	Symbols     *[]string        `json:"symbols"`
}

func createSession(w http.ResponseWriter, r *http.Request) {
	defer must.Recover()

	var request CreateSessionRequest

	must.Succeed(json.NewDecoder(r.Body).Decode(&request)).ElseRespond(w, http.StatusBadRequest)

	session := must.Return(session.New(request.SessionName, *request.Symbols)).ElseRespond(w, http.StatusBadRequest)

	store.Instance().Add(session)

	json := must.Return(json.Marshal(&session)).ElseRespond(w, http.StatusBadRequest)

	w.WriteHeader(http.StatusOK)
	w.Write(json)
}
