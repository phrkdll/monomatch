package session

import (
	"log/slog"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/phrkdll/monomatch/pkg/player"
	"github.com/phrkdll/monomatch/pkg/session"
	"github.com/phrkdll/monomatch/pkg/session/store"
	"github.com/phrkdll/must/pkg/must"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins for simplicity; adjust as needed for security
	},
}

func sessionSocket(w http.ResponseWriter, r *http.Request) {
	defer must.Recover()

	id := r.URL.Query().Get("sessionId")

	s := must.Return(store.Instance().Get(session.SessionId{Inner: id})).ElseRespond(w, http.StatusBadRequest)

	// upgrade this connection to a WebSocket
	// connection
	conn := must.Return(upgrader.Upgrade(w, r, nil)).ElseRespond(w, http.StatusBadRequest)

	// listen indefinitely for new messages coming
	// through on our WebSocket connection
	for {
		// read in a message
		var message any
		must.Succeed(conn.ReadJSON(&message)).ElseRespond(w, http.StatusBadRequest)

		// print out that message for clarity
		slog.Info("message received", "message", message)
		state := session.SessionState{
			Id:             s.Id,
			Name:           s.Name,
			CreatedAt:      s.CreatedAt,
			TopMostCard:    must.Return(s.Cards.Top()).ElsePanic(),
			RemainingCards: s.Cards.Len(),
			Players:        []player.PlayerState{},
		}

		must.Succeed(conn.WriteJSON(&state)).ElseRespond(w, http.StatusBadRequest)
	}
}
