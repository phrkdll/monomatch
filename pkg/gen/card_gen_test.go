package gen_test

import (
	"slices"
	"testing"

	"github.com/phrkdll/monomatch/internal/testdata"
	"github.com/phrkdll/monomatch/pkg/card"
	"github.com/phrkdll/monomatch/pkg/gen"
	"github.com/stretchr/testify/assert"
)

func TestGenerateCards(t *testing.T) {
	type testCase struct {
		name      string
		symbols   []string
		expectErr error
	}

	testCases := []testCase{
		{
			name:      "No symbols provided",
			symbols:   []string{},
			expectErr: gen.ErrNotEnoughSymbolsProvided,
		},
		{
			name:      "Not enough symbols provided",
			symbols:   []string{"Lonely symbol"},
			expectErr: gen.ErrNotEnoughSymbolsProvided,
		},
		{
			name:      "Exact number of symbols provided",
			symbols:   testdata.SymbolNames,
			expectErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cards, err := gen.GenerateCards(gen.MakeSymbols(tc.symbols))

			assert.Equal(t, tc.expectErr, err)
			if err == nil {
				assert.Len(t, cards, len(tc.symbols))

				ids := []card.CardId{}
				var prevCard *card.Card
				for _, card := range cards {
					if !assert.False(t, slices.Contains(ids, card.ID)) || !assert.Len(t, card.Symbols, 8) {
						return
					}
					ids = append(ids, card.ID)
					if prevCard == nil {
						continue
					}

					combined := append(prevCard.Symbols, card.Symbols...)
					compact := slices.Compact(combined)

					assert.Len(t, compact, 15)

					prevCard = &card
				}
			}
		})
	}
}
