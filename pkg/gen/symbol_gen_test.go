package gen_test

import (
	"slices"
	"testing"

	"github.com/phrkdll/monomatch/internal/testdata"
	"github.com/phrkdll/monomatch/pkg/gen"
	"github.com/phrkdll/monomatch/pkg/symbol"
	"github.com/stretchr/testify/assert"
)

func TestMakeSymbols(t *testing.T) {
	symbols := gen.MakeSymbols(testdata.SymbolNames)

	assert.Len(t, symbols, len(testdata.SymbolNames))

	var symbolIds []symbol.SymbolId
	for _, r := range symbols {
		symbolIds = append(symbolIds, r.Id)
	}

	assert.Len(t, slices.Compact(symbolIds), len(symbols))
}
