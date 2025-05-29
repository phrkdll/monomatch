package gen

import (
	"github.com/google/uuid"
	"github.com/phrkdll/monomatch/pkg/symbol"
)

func MakeSymbols(names []string) []symbol.Symbol {
	var symbols []symbol.Symbol

	for _, name := range names {
		symbols = append(symbols, symbol.Symbol{
			ID:   symbol.SymbolId{Inner: uuid.New().String()},
			Name: name,
		})
	}

	return symbols
}
