package session_test

import (
	"testing"

	"github.com/phrkdll/monomatch/internal/testdata"
	"github.com/phrkdll/monomatch/pkg/gen"
	"github.com/phrkdll/monomatch/pkg/player"
	"github.com/phrkdll/monomatch/pkg/session"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	type testCase struct {
		name    string
		symbols []string
		err     error
	}

	testCases := []testCase{
		{
			name:    "complete set of symbols",
			symbols: testdata.SymbolNames,
			err:     nil,
		},
		{
			name:    "insufficient set of symbols",
			symbols: []string{"only 1"},
			err:     gen.ErrNotEnoughSymbolsProvided,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testSession, err := session.New(tc.name, tc.symbols)

			assert.Equal(t, tc.err, err)

			if err == nil {
				assert.NotNil(t, testSession)
			}
		})
	}
}

func TestAddPlayer(t *testing.T) {
	type testCase struct {
		name       string
		playerName string
		err        error
	}

	testCases := []testCase{
		{
			name:       "add succeeds",
			playerName: "test",
			err:        nil,
		},
		{
			name:       "add fails",
			playerName: "test",
			err:        session.ErrPlayerNameAlreadyTaken,
		},
		{
			name:       "add fails",
			playerName: "",
			err:        player.ErrPlayerNameRequired,
		},
	}

	testSession, _ := session.New("test", testdata.SymbolNames)
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			player, err := testSession.AddPlayer(tc.playerName)
			assert.Equal(t, tc.err, err)

			if err == nil {
				assert.Equal(t, tc.playerName, player.Name)
			}
		})
	}
}
