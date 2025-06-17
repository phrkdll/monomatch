package card

import (
	"errors"

	"github.com/google/uuid"
	"github.com/phrkdll/monomatch/pkg/symbol"
	"github.com/phrkdll/strongoid/pkg/strongoid"
)

//go:generate go run github.com/phrkdll/strongoid/cmd/gen --modules=json,
type CardId strongoid.Id[string]

type Card struct {
	Id      CardId          `json:"id"`
	Symbols []symbol.Symbol `json:"symbols"`
}

var (
	ErrInsufficientSymbols = errors.New("insufficient symbols provided")
	ErrTooManySymbols      = errors.New("too many symbols provided")
)

func New(symbols []symbol.Symbol) (*Card, error) {
	if len(symbols) < 8 {
		return nil, ErrInsufficientSymbols
	}

	if len(symbols) > 8 {
		return nil, ErrTooManySymbols
	}

	return &Card{
		Id:      CardId{Inner: uuid.NewString()},
		Symbols: symbols,
	}, nil
}
