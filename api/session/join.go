package session

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/phrkdll/monomatch/internal/utils"
	"github.com/phrkdll/monomatch/pkg/session"
	"github.com/phrkdll/monomatch/pkg/session/store"
	"github.com/phrkdll/must/pkg/must"
)

type JoinRequest struct {
	PlayerName string `json:"playerName"`
}

func joinSession(w http.ResponseWriter, r *http.Request) {
	defer must.Recover()

	sessionId := chi.URLParam(r, "sessionId")

	var request JoinRequest

	must.Succeed(json.NewDecoder(r.Body).Decode(&request)).
		ElseRespond(w, http.StatusBadRequest)

	session := must.Return(store.Instance().Get(session.SessionId{Inner: sessionId})).
		ElseRespond(w, http.StatusBadRequest)

	p := must.Return(session.AddPlayer(request.PlayerName)).
		ElseRespond(w, http.StatusBadRequest)

	utils.SendJSON(w, http.StatusOK, p)
}
