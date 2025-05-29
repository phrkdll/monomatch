package player

import (
	"errors"

	"github.com/google/uuid"
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
	ID    PlayerId               `json:"id"`
	Name  string                 `json:"name"`
	Cards stack.Stack[card.Card] `json:"-"`
}

func New(name string) (*Player, error) {
	if name == "" {
		return nil, ErrPlayerNameRequired
	}

	return &Player{
		ID:   PlayerId{Inner: uuid.NewString()},
		Name: name,
	}, nil
}
