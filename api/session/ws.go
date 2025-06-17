package session

import (
	"encoding/json"
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

	sendSessionState(w, conn, s)

	// listen indefinitely for new messages coming
	// through on our WebSocket connection
	for {
		// read in a message
		var message map[string]any
		must.Succeed(conn.ReadJSON(&message)).ElseRespond(w, http.StatusBadRequest)

		jsonString := must.Return(json.Marshal(&message)).ElseRespond(w, http.StatusBadRequest)

		mt, ok := message["messageType"].(string)
		if !ok {
			break
		}

		switch mt {
		case "add-player":
			var r AddPlayerRequest
			must.Succeed(json.Unmarshal(jsonString, &r))
			if err := s.AddPlayer(r.Id, r.Name, conn); err != nil {
				continue
			}
		}

		// print out that message for clarity
		slog.Info("message received", "message", message)
		sendSessionState(w, conn, s)
	}
}

func sendSessionState(w http.ResponseWriter, conn *websocket.Conn, s *session.Session) {
	state := session.SessionState{
		Id:             s.Id,
		Name:           s.Name,
		CreatedAt:      s.CreatedAt,
		TopMostCard:    must.Return(s.Cards.Top()).ElsePanic(),
		RemainingCards: s.Cards.Len(),
	}

	for _, p := range s.Players {
		state.Players = append(state.Players, player.PlayerState{
			PlayerName:  p.Name,
			TopMostCard: must.Return(p.Cards.Top()).ElsePanic(),
			CardCount:   p.Cards.Len(),
		})
	}

	must.Succeed(conn.WriteJSON(&state)).ElseRespond(w, http.StatusBadRequest)
}

type AddPlayerRequest struct {
	Id   player.PlayerId `json:"id"`
	Name string          `json:"Name"`
}
