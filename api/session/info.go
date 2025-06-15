package session

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/phrkdll/monomatch/internal/utils"
	"github.com/phrkdll/monomatch/pkg/session"
	"github.com/phrkdll/monomatch/pkg/session/store"
	"github.com/phrkdll/must/pkg/must"
)

var (
	ErrSessionIdRequired = errors.New("session id is required")
)

func getSessionInfo(w http.ResponseWriter, r *http.Request) {
	defer must.Recover()

	sessionId := chi.URLParam(r, "sessionId")

	if sessionId == "" || sessionId == "undefined" {
		must.Succeed(ErrSessionIdRequired).ElseRespond(w, http.StatusBadRequest)
	}

	session := must.Return(store.Instance().Get(session.SessionId{Inner: sessionId})).ElseRespond(w, http.StatusNotFound)

	utils.SendJSON(w, http.StatusOK, session)
}
