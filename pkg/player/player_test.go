package player_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/phrkdll/monomatch/pkg/player"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	type testCase struct {
		name     string
		playerId player.PlayerId
		player   string
		err      error
	}

	testCases := []testCase{
		{
			name:     "valid",
			playerId: player.PlayerId{Inner: uuid.NewString()},
			player:   "vlad",
			err:      nil,
		},
		{
			name:     "no name",
			playerId: player.PlayerId{Inner: uuid.NewString()},
			player:   "",
			err:      player.ErrPlayerNameRequired,
		},
		{
			name:     "no id",
			playerId: player.PlayerId{},
			player:   "bork",
			err:      player.ErrPlayerNameRequired,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			p, err := player.New(tc.playerId, tc.player, nil)
			assert.Equal(t, tc.err, err)

			if err == nil {
				assert.Equal(t, tc.player, p.Name)
				assert.Equal(t, 0, p.Cards.Len())
				assert.NotNil(t, p.Id.Inner)
			}
		})
	}
}
