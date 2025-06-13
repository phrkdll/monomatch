package symbol_test

import (
	"testing"

	"github.com/phrkdll/monomatch/pkg/symbol"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	type testCase struct {
		name   string
		symbol string
		err    error
	}

	testCases := []testCase{
		{
			name:   "valid symbol",
			symbol: "sym",
			err:    nil,
		},
		{
			name:   "invalid symbol",
			symbol: "",
			err:    symbol.ErrSymbolNameRequired,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s, err := symbol.New(tc.symbol)
			assert.Equal(t, tc.err, err)

			if err == nil {
				assert.Equal(t, tc.symbol, s.Name)
				assert.NotNil(t, s.Id.Inner)
			}
		})
	}
}
