package symbol

import (
	"errors"

	"github.com/google/uuid"
	"github.com/phrkdll/strongoid/pkg/strongoid"
)

//go:generate go run github.com/phrkdll/strongoid/cmd/gen --modules=json,
type SymbolId strongoid.Id[string]

type Symbol struct {
	Id   SymbolId `json:"id"`
	Name string   `json:"name"`
}

var (
	ErrSymbolNameRequired = errors.New("a symbol name is required")
)

func New(name string) (*Symbol, error) {
	if name == "" {
		return nil, ErrSymbolNameRequired
	}

	return &Symbol{
		Id:   SymbolId{Inner: uuid.NewString()},
		Name: name,
	}, nil
}
