package gen

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/phrkdll/monomatch/pkg/symbol"
)

func MakeSymbols(names []string) []symbol.Symbol {
	var symbols []symbol.Symbol

	rand.New(rand.NewSource(time.Now().UnixNano()))
	indices := rand.Perm(len(names))

	var loop int
	for _, index := range indices {
		if loop == 57 {
			break
		}

		symbols = append(symbols, symbol.Symbol{
			ID:   symbol.SymbolId{Inner: uuid.New().String()},
			Name: names[index],
		})

		loop++
	}

	return symbols
}
