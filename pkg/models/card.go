package models

import "github.com/phrkdll/strongoid/pkg/strongoid"

type CardId strongoid.Id[string]

type Card struct {
	ID      CardId
	Symbols []Symbol
}

func NewCardId(inner string) CardId {
	return CardId{Inner: inner}
}
