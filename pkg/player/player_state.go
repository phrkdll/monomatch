package player

import "github.com/phrkdll/monomatch/pkg/card"

type PlayerState struct {
	PlayerName  string     `json:"playerName"`
	TopMostCard *card.Card `json:"topMostCard"`
	CardCount   int        `json:"cardCount"`
}
