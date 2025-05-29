package gen

import (
	"errors"

	"github.com/phrkdll/monomatch/pkg/card"
	"github.com/phrkdll/monomatch/pkg/symbol"
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
	cards = append(cards, *card.New(symbols[:n+1]))

	// 2. n groups with n cards each
	for i := range n {
		for j := range n {
			card := card.New([]symbol.Symbol{symbols[i+1]})
			for k := range n {
				index := n + 1 + n*k + ((i*k + j) % n)
				card.Symbols = append(card.Symbols, symbols[index])
			}
			cards = append(cards, *card)
		}
	}

	// 3. Last n cards
	for k := range n {
		card := card.New([]symbol.Symbol{symbols[0]})
		for i := range n {
			index := n + 1 + n*k + i
			card.Symbols = append(card.Symbols, symbols[index])
		}
		cards = append(cards, *card)
	}

	return cards, nil
}
