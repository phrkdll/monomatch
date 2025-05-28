package gen

import (
	"slices"
	"testing"

	"github.com/phrkdll/monomatch/pkg/models"
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
			expectErr: ErrNotEnoughSymbolsProvided,
		},
		{
			name:      "Not enough symbols provided",
			symbols:   []string{"Lonely symbol"},
			expectErr: ErrNotEnoughSymbolsProvided,
		},
		{
			name:      "Exact number of symbols provided",
			symbols:   testSymbols,
			expectErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cards, err := GenerateCards(MakeSymbols(tc.symbols))

			assert.Equal(t, tc.expectErr, err)
			if err == nil {
				assert.Len(t, cards, len(tc.symbols))

				ids := []models.CardId{}
				for _, card := range cards {
					if assert.False(t, slices.Contains(ids, card.ID)) {
						return
					}
					ids = append(ids, card.ID)
				}
			}
		})
	}

}
