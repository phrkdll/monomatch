package models

import "github.com/phrkdll/strongoid/pkg/strongoid"

type SymbolId strongoid.Id[int64]

type Symbol struct {
	ID   SymbolId
	Name string
}

func NewSymbolId(inner int64) SymbolId {
	return SymbolId{Inner: inner}
}
