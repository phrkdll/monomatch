package gen

import "github.com/phrkdll/monomatch/pkg/models"

func MakeSymbols(names []string) []models.Symbol {
	symbols := make([]models.Symbol, len(names))

	for i, name := range names {
		symbols[i] = models.Symbol{
			ID:   models.NewSymbolId(int64(i)),
			Name: name,
		}
	}

	return symbols
}
