package player_test

import (
	"testing"

	"github.com/phrkdll/monomatch/pkg/player"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	type testCase struct {
		name   string
		player string
		err    error
	}

	testCases := []testCase{
		{
			name:   "valid",
			player: "vlad",
			err:    nil,
		},
		{
			name:   "invalid",
			player: "",
			err:    player.ErrPlayerNameRequired,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			p, err := player.New(tc.player)
			assert.Equal(t, tc.err, err)

			if err == nil {
				assert.Equal(t, tc.player, p.Name)
				assert.Equal(t, 0, p.Cards.Len())
				assert.NotNil(t, p.ID.Inner)
			}
		})
	}
}
