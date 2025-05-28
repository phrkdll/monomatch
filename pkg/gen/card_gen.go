package gen

import (
	"errors"

	"github.com/google/uuid"
	"github.com/phrkdll/monomatch/pkg/models"
)

const n = 7 // Ordnung der projektiven Ebene

var (
	ErrNotEnoughSymbolsProvided = errors.New("insufficient symbols provided")
)

func GenerateCards(symbols []models.Symbol) ([]models.Card, error) {
	totalSymbols := n*n + n + 1
	if len(symbols) != totalSymbols {
		return []models.Card{}, ErrNotEnoughSymbolsProvided
	}

	var cards []models.Card
	cardId := models.NewCardId(getUuid())

	// 1. Erste Karte
	cards = append(cards, models.Card{
		ID:      cardId,
		Symbols: symbols[:n+1],
	})
	cardId = models.NewCardId(getUuid())

	// 2. n Gruppen mit je n Karten
	for i := range n {
		for j := range n {
			card := models.Card{
				ID:      cardId,
				Symbols: []models.Symbol{symbols[i+1]},
			}
			for k := range n {
				index := n + 1 + n*k + ((i*k + j) % n)
				card.Symbols = append(card.Symbols, symbols[index])
			}
			cards = append(cards, card)
			cardId = models.NewCardId(getUuid())
		}
	}

	// 3. Letzte n Karten
	for k := range n {
		card := models.Card{
			ID:      cardId,
			Symbols: []models.Symbol{symbols[0]},
		}
		for i := range n {
			index := n + 1 + n*k + i
			card.Symbols = append(card.Symbols, symbols[index])
		}
		cards = append(cards, card)
		cardId = models.NewCardId(getUuid())
	}

	return cards, nil
}

func getUuid() string {
	id, _ := uuid.NewRandom()

	return id.String()
}
