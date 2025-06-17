package player

import (
	"errors"

	"github.com/gorilla/websocket"
	"github.com/phrkdll/monomatch/pkg/card"
	"github.com/phrkdll/monomatch/pkg/stack"
	"github.com/phrkdll/strongoid/pkg/strongoid"
)

var (
	ErrIdIsRequired       = errors.New("a player id is required")
	ErrPlayerNameRequired = errors.New("a player name is required")
)

//go:generate go run github.com/phrkdll/strongoid/cmd/gen --modules=json,
type PlayerId strongoid.Id[string]

type Player struct {
	Id    PlayerId
	Name  string
	Cards stack.Stack[card.Card]
	Conn  *websocket.Conn
	Ready bool
}

func New(id PlayerId, name string, conn *websocket.Conn) (*Player, error) {
	if id.Inner == "" {
		return nil, ErrIdIsRequired
	}

	if name == "" {
		return nil, ErrPlayerNameRequired
	}

	return &Player{
		Id:    id,
		Name:  name,
		Cards: stack.Stack[card.Card]{},
		Conn:  conn,
	}, nil
}
