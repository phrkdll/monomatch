package gen

import (
	"errors"

	"github.com/phrkdll/monomatch/pkg/card"
	"github.com/phrkdll/monomatch/pkg/symbol"
	"github.com/phrkdll/must/pkg/must"
)

const n = 7 // Order of the projective plane

var (
	ErrNotEnoughSymbolsProvided = errors.New("insufficient symbols provided")
)

func GenerateCards(symbols []symbol.Symbol) ([]card.Card, error) {
	totalSymbols := n*n + n + 1
	if len(symbols) != totalSymbols {
		return []card.Card{}, ErrNotEnoughSymbolsProvided
	}

	var cards []card.Card

	// 1. First card
	cards = append(cards, *must.Return(card.New(symbols[:n+1])).ElsePanic())

	// 2. n groups with n cards each
	for i := range n {
		for j := range n {
			groupSymbols := []symbol.Symbol{symbols[i+1]}
			for k := range n {
				index := n + 1 + n*k + ((i*k + j) % n)
				groupSymbols = append(groupSymbols, symbols[index])
			}

			cards = append(cards, *must.Return(card.New(groupSymbols)).ElsePanic())
		}
	}

	// 3. Last n cards
	for k := range n {
		lastSymbols := []symbol.Symbol{symbols[0]}
		for i := range n {
			index := n + 1 + n*k + i
			lastSymbols = append(lastSymbols, symbols[index])
		}

		cards = append(cards, *must.Return(card.New(lastSymbols)).ElsePanic())
	}

	return cards, nil
}
