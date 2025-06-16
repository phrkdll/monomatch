package session

import (
	"time"

	"github.com/phrkdll/monomatch/pkg/card"
	"github.com/phrkdll/monomatch/pkg/player"
)

type SessionState struct {
	Id             SessionId            `json:"id"`
	Name           string               `json:"name"`
	CreatedAt      time.Time            `json:"createdAt"`
	TopMostCard    *card.Card           `json:"topMostCard"`
	RemainingCards int                  `json:"remainingCards"`
	Players        []player.PlayerState `json:"players"`
}
