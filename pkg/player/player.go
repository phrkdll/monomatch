package player

import (
	"errors"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/phrkdll/monomatch/pkg/card"
	"github.com/phrkdll/monomatch/pkg/stack"
	"github.com/phrkdll/strongoid/pkg/strongoid"
)

var (
	ErrPlayerNameRequired = errors.New("a player name is required")
)

//go:generate go run github.com/phrkdll/strongoid/cmd/gen --modules=json,
type PlayerId strongoid.Id[string]

type Player struct {
	Id    PlayerId
	Name  string
	Cards stack.Stack[card.Card]
	Conn  *websocket.Conn
}

func New(name string, conn *websocket.Conn) (*Player, error) {
	if name == "" {
		return nil, ErrPlayerNameRequired
	}

	return &Player{
		Id:    PlayerId{Inner: uuid.NewString()},
		Name:  name,
		Cards: stack.Stack[card.Card]{},
		Conn:  conn,
	}, nil
}
