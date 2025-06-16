package session

import (
	"errors"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/phrkdll/monomatch/pkg/card"
	"github.com/phrkdll/monomatch/pkg/gen"
	"github.com/phrkdll/monomatch/pkg/player"
	"github.com/phrkdll/monomatch/pkg/stack"
	"github.com/phrkdll/strongoid/pkg/strongoid"
)

var (
	ErrPlayerNameAlreadyTaken = errors.New("player name already taken")
)

//go:generate go run github.com/phrkdll/strongoid/cmd/gen --modules=json,
type SessionId strongoid.Id[string]

type Session struct {
	Id        SessionId
	Name      string
	CreatedAt time.Time
	Cards     stack.Stack[card.Card]
	Players   []player.Player
}

func New(name string, input []string) (*Session, error) {
	symbols := gen.MakeSymbols(input)
	cards, err := gen.GenerateCards(symbols)

	stack := stack.New([]card.Card{})
	rand.New(rand.NewSource(time.Now().UnixNano()))
	indices := rand.Perm(len(cards))

	var loop int
	for _, index := range indices {
		if loop == len(cards) {
			break
		}

		stack.Push(cards[index])

		loop++
	}

	if err != nil {
		return nil, err
	}

	return &Session{
		Id:        SessionId{Inner: uuid.New().String()},
		Name:      name,
		CreatedAt: time.Now().UTC(),
		Cards:     stack,
		Players:   []player.Player{},
	}, nil
}

func (s *Session) AddPlayer(name string, conn *websocket.Conn) (*player.Player, error) {
	for _, p := range s.Players {
		if p.Name == name {
			return nil, ErrPlayerNameAlreadyTaken
		}
	}

	player, err := player.New(name, conn)
	if err != nil {
		return nil, err
	}

	s.Players = append(s.Players, *player)

	return player, nil
}
