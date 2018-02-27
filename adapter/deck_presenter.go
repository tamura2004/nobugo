package adapter

import (
	"github.com/tamura2004/nobugo/domain"
)

type DeckPresenter domain.Deck

func (p DeckPresenter) Names() []string {
	names := []string{}
	for _, card := range p.Card {
		names = append(names, card.Name)
	}
	return names
}
