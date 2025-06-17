package player

import "github.com/phrkdll/monomatch/pkg/card"

type PlayerState struct {
	Id          PlayerId   `json:"id"`
	PlayerName  string     `json:"playerName"`
	TopMostCard *card.Card `json:"topMostCard"`
	CardCount   int        `json:"cardCount"`
	Ready       bool       `json:"ready"`
}
