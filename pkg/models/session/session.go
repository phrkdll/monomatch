package session

import (
	"github.com/google/uuid"
	"github.com/phrkdll/monomatch/pkg/gen"
	"github.com/phrkdll/monomatch/pkg/models/card"
	"github.com/phrkdll/strongoid/pkg/strongoid"
)

//go:generate go run github.com/phrkdll/strongoid/cmd/gen --modules=json,
type SessionId strongoid.Id[string]

type Session struct {
	ID    SessionId   `json:"id"`
	Cards []card.Card `json:"cards"`
}

func New(input []string) (*Session, error) {
	symbols := gen.MakeSymbols(input)
	cards, err := gen.GenerateCards(symbols)
	if err != nil {
		return nil, err
	}

	return &Session{
		ID:    SessionId{Inner: uuid.New().String()},
		Cards: cards,
	}, nil
}
