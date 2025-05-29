package card

import (
	"github.com/google/uuid"
	"github.com/phrkdll/monomatch/pkg/symbol"
	"github.com/phrkdll/strongoid/pkg/strongoid"
)

//go:generate go run github.com/phrkdll/strongoid/cmd/gen --modules=json,
type CardId strongoid.Id[string]

type Card struct {
	ID      CardId          `json:"id"`
	Symbols []symbol.Symbol `json:"symbols"`
}

func New(symbols []symbol.Symbol) *Card {
	return &Card{
		ID:      CardId{Inner: uuid.New().String()},
		Symbols: symbols,
	}
}
