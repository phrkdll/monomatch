package symbol

import "github.com/phrkdll/strongoid/pkg/strongoid"

//go:generate go run github.com/phrkdll/strongoid/cmd/gen --modules=json,
type SymbolId strongoid.Id[string]

type Symbol struct {
	ID   SymbolId `json:"id"`
	Name string   `json:"name"`
}
