package session

import (
	"encoding/json"
	"net/http"

	"github.com/phrkdll/monomatch/internal/utils"
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

	must.Succeed(json.NewDecoder(r.Body).Decode(&request)).ElseRespond(w, http.StatusBadRequest)

	session := must.Return(store.Instance().Get(request.Id)).ElseRespond(w, http.StatusBadRequest)

	p := must.Return(session.AddPlayer(request.PlayerName)).ElseRespond(w, http.StatusBadRequest)

	json := must.Return(json.Marshal(&p)).ElseRespond(w, http.StatusBadRequest)

	utils.SendJSON(w, http.StatusOK, json)
}
