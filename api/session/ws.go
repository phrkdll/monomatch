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

	sendSessionState(w, []*websocket.Conn{conn}, s)

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
		case "player-ready":
			var r PlayerReadyRequest
			must.Succeed(json.Unmarshal(jsonString, &r))
			s.Players[r.Id].Ready = r.Ready
		}

		// print out that message for clarity
		slog.Info("message received", "message", message)

		var conns []*websocket.Conn
		for _, player := range s.Players {
			conns = append(conns, player.Conn)
		}

		sendSessionState(w, conns, s)
	}
}

func sendSessionState(w http.ResponseWriter, conns []*websocket.Conn, s *session.Session) {
	var allReady bool
	for _, player := range s.Players {
		allReady = player.Ready
		if allReady == false {
			break
		}
	}

	state := session.SessionState{
		Id:             s.Id,
		Name:           s.Name,
		CreatedAt:      s.CreatedAt,
		TopMostCard:    valueOrNil(allReady, must.Return(s.Cards.Top()).ElsePanic()),
		RemainingCards: s.Cards.Len(),
	}

	for _, p := range s.Players {
		state.Players = append(state.Players, player.PlayerState{
			Id:          p.Id,
			PlayerName:  p.Name,
			TopMostCard: valueOrNil(allReady, must.Return(p.Cards.Top()).ElsePanic()),
			CardCount:   p.Cards.Len(),
			Ready:       p.Ready,
		})
	}

	for _, conn := range conns {
		must.Succeed(conn.WriteJSON(&state)).ElseRespond(w, http.StatusBadRequest)
	}
}

func valueOrNil[T any](predicate bool, value T) (ret T) {
	if predicate {
		ret = value
		return
	}
	return
}

type AddPlayerRequest struct {
	Id   player.PlayerId `json:"id"`
	Name string          `json:"name"`
}

type PlayerReadyRequest struct {
	Id    player.PlayerId `json:"id"`
	Ready bool            `json:"ready"`
}
