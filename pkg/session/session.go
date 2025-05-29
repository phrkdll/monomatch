package session

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/phrkdll/monomatch/pkg/card"
	"github.com/phrkdll/monomatch/pkg/gen"
	"github.com/phrkdll/monomatch/pkg/player"
	"github.com/phrkdll/strongoid/pkg/strongoid"
)

var (
	ErrPlayerNameAlreadyTaken = errors.New("player name already taken")
)

//go:generate go run github.com/phrkdll/strongoid/cmd/gen --modules=json,
type SessionId strongoid.Id[string]

type Session struct {
	ID        SessionId       `json:"id"`
	CreatedAt time.Time       `json:"createdAt"`
	Cards     []card.Card     `json:"-"`
	Players   []player.Player `json:"-"`
}

func New(input []string) (*Session, error) {
	symbols := gen.MakeSymbols(input)
	cards, err := gen.GenerateCards(symbols)
	if err != nil {
		return nil, err
	}

	return &Session{
		ID:        SessionId{Inner: uuid.New().String()},
		CreatedAt: time.Now().UTC(),
		Cards:     cards,
	}, nil
}

func (s *Session) NewPlayer(name string) (*player.Player, error) {
	for _, p := range s.Players {
		if p.Name == name {
			return nil, ErrPlayerNameAlreadyTaken
		}
	}

	player, err := player.New(name)
	if err != nil {
		return nil, err
	}

	s.Players = append(s.Players, *player)

	return player, nil
}
