package card_test

import (
	"testing"

	"github.com/phrkdll/monomatch/internal/testdata"
	"github.com/phrkdll/monomatch/pkg/card"
	"github.com/phrkdll/monomatch/pkg/gen"
	"github.com/phrkdll/monomatch/pkg/symbol"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	type testCase struct {
		name    string
		symbols []symbol.Symbol
		err     error
	}

	testCases := []testCase{
		{
			name:    "valid",
			symbols: testdata.EightSymbols,
			err:     nil,
		},
		{
			name:    "insufficient symbols",
			symbols: gen.MakeSymbols([]string{"00"}),
			err:     card.ErrInsufficientSymbols,
		},
		{
			name:    "too many symbols",
			symbols: append(testdata.EightSymbols, gen.MakeSymbols([]string{"09"})[0]),
			err:     card.ErrTooManySymbols,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c, err := card.New(tc.symbols)
			assert.Equal(t, tc.err, err)

			if err == nil {
				assert.Equal(t, tc.symbols, c.Symbols)
				assert.NotNil(t, c.Id.Inner)
			}
		})
	}
}
